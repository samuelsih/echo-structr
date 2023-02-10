package presentation

import (
	"github.com/labstack/echo/v4"
	"github.com/samuelsih/echo-structr/business"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	presentationServer = echo.New()
)

type P struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestParseInput(t *testing.T) {
	inputJSON := `{"name":"tester","age":23}`
	var input P

	t.Run("with fulfilled cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(inputJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.AddCookie(&http.Cookie{
			Name:   cookieName,
			Value:  "some-token",
			MaxAge: cookieExpired,
		})

		rec := httptest.NewRecorder()

		cr := business.CommonRequest{}

		c := presentationServer.NewContext(req, rec)

		err := ParseInput(c, &input, &cr)

		if err != nil {
			log.Fatal(err)
		}

		Equal(t, 200, rec.Code)
		Equal(t, input.Name, "tester")
		Equal(t, input.Age, 23)
		Equal(t, cr.Token, "some-token")
	})

	t.Run("empty cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(inputJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		cr := business.CommonRequest{}

		c := presentationServer.NewContext(req, rec)

		err := ParseInput(c, &input, &cr)

		if err != nil {
			log.Fatal(err)
		}

		Equal(t, 200, rec.Code)
		Equal(t, input.Name, "tester")
		Equal(t, input.Age, 23)
		Equal(t, cr.Token, "")
	})
}

func TestRenderOutput(t *testing.T) {
	t.Parallel()

	type Foo struct {
		cr business.CommonResponse
		p  P
	}

	t.Run("with set token", func(t *testing.T) {
		foo := Foo{
			cr: business.CommonResponse{
				SetAuthToken: "123123123",
				StatusCode:   0,
				Msg:          "OK",
			},
			p: P{
				Name: "asep",
				Age:  23,
			},
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := presentationServer.NewContext(req, rec)

		err := RenderOutput(c, &foo, foo.cr)

		NoError(t, err)
		Equal(t, 200, rec.Code)
		NotEmpty(t, rec.Header().Get("Set-Cookie"))
	})

	t.Run("without set token", func(t *testing.T) {
		foo := Foo{
			cr: business.CommonResponse{
				StatusCode: 0,
				Msg:        "OK",
			},
			p: P{
				Name: "asep",
				Age:  23,
			},
		}

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := presentationServer.NewContext(req, rec)

		err := RenderOutput(c, &foo, foo.cr)

		NoError(t, err)
		Equal(t, 200, rec.Code)
	})
}

func Equal[T comparable](t *testing.T, a, b T) {
	t.Helper()

	if a != b {
		t.Fatalf("expected %v, got %v", a, b)
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected not nil error, got %v", err)
	}
}

func NotEmpty[T comparable](t *testing.T, a T) {
	t.Helper()

	objValue := reflect.ValueOf(a)

	switch objValue.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice:
		if objValue.Len() == 0 {
			t.Fatalf("item is empty")
		}

	default:
		zero := reflect.Zero(objValue.Type())
		if reflect.DeepEqual(a, zero.Interface()) {
			t.Fatalf("item is empty")
		}

	}
}
