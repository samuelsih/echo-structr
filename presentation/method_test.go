package presentation

import (
	"context"
	"github.com/labstack/echo/v4"
	b "github.com/samuelsih/echo-structr/business"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	methodServer = echo.New()
)

func TestGET(t *testing.T) {
	t.Parallel()

	type A struct {
		b.CommonResponse
		Name      string `json:"name,omitempty"`
		Something string `json:"something,omitempty"`
	}

	var f GetFunc[A] = func(ctx context.Context, cr b.CommonRequest) A {
		if cr.Token == "" {
			t.Fatalf("cookie does not exist")
		}

		var a A
		a.Name = "tester"
		a.Something = "13123123"
		return a
	}

	handler := GET(f, Default)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.AddCookie(&http.Cookie{
		Name:   cookieName,
		Value:  "some-token",
		MaxAge: cookieExpired,
	})

	rec := httptest.NewRecorder()

	c := methodServer.NewContext(req, rec)

	NoError(t, handler(c))
	Equal(t, 200, rec.Code)
}

func TestGETWithOpts(t *testing.T) {
	t.Parallel()

	type A struct {
		b.CommonResponse
		Name      string `json:"name,omitempty"`
		Something string `json:"something,omitempty"`
	}

	t.Run("query param", func(t *testing.T) {
		var f GetFunc[A] = func(ctx context.Context, cr b.CommonRequest) A {
			if _, ok := cr.QueryParams["param1"]; !ok {
				t.Fatalf("query param (param1) does not exist")
			}

			if _, ok := cr.QueryParams["param2"]; !ok {
				t.Fatalf("query param (param2) does not exist")
			}

			var a A
			a.Name = "tester"
			a.Something = "13123123"
			return a
		}

		handler := GET(f, With(QueryParams("param1", "param2")))

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := methodServer.NewContext(req, rec)

		NoError(t, handler(c))
		Equal(t, 200, rec.Code)
	})

	t.Run("url param", func(t *testing.T) {
		var f GetFunc[A] = func(ctx context.Context, cr b.CommonRequest) A {
			if _, ok := cr.URLParams["param1"]; !ok {
				t.Fatalf("url param (param1) does not exist")
			}

			if _, ok := cr.URLParams["param2"]; !ok {
				t.Fatalf("url param (param2) does not exist")
			}

			var a A
			a.Name = "tester"
			a.Something = "13123123"
			return a
		}

		handler := GET(f, With(UrlParams("param1", "param2")))

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := methodServer.NewContext(req, rec)

		NoError(t, handler(c))
		Equal(t, 200, rec.Code)
	})
}

func TestPOST(t *testing.T) {
	t.Parallel()

	type In struct {
		b.CommonRequest
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	type Out struct {
		b.CommonResponse
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	var f PostFunc[In, Out] = func(ctx context.Context, in In, cr b.CommonRequest) Out {
		if cr.Token == "" {
			t.Fatalf("cookie does not exist")
		}

		if in.Name != "tester" {
			t.Fatalf("name must be tester, got %s", in.Name)
		}

		if in.Age != 23 {
			t.Fatalf("age must be 23, got %d", in.Age)
		}

		var a Out

		a.Name = "tester"
		a.Age = 23
		return a
	}

	handler := POST(f, Default)
	input := `{"name":"tester","age":23}`

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(input))
	req.AddCookie(&http.Cookie{
		Name:   cookieName,
		Value:  "some-token",
		MaxAge: cookieExpired,
	})
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := methodServer.NewContext(req, rec)

	NoError(t, handler(c))
	Equal(t, 200, rec.Code)
}

func TestPOSTWithOpts(t *testing.T) {
	t.Parallel()

	type In struct {
		b.CommonRequest
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	type Out struct {
		b.CommonResponse
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	t.Run("url param", func(t *testing.T) {
		var f PostFunc[In, Out] = func(ctx context.Context, in In, cr b.CommonRequest) Out {
			if _, ok := cr.URLParams["param1"]; !ok {
				t.Fatalf("url param (param1) does not exist")
			}

			if _, ok := cr.URLParams["param2"]; !ok {
				t.Fatalf("url param (param2) does not exist")
			}

			if in.Name != "tester" {
				t.Fatalf("name must be tester, got %s", in.Name)
			}

			if in.Age != 23 {
				t.Fatalf("age must be 23, got %d", in.Age)
			}

			var a Out

			a.Name = "tester"
			a.Age = 23
			return a
		}

		handler := POST(f, With(UrlParams("param1", "param2")))
		input := `{"name":"tester","age":23}`

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(input))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := methodServer.NewContext(req, rec)

		NoError(t, handler(c))
		Equal(t, 200, rec.Code)
	})

	t.Run("query param", func(t *testing.T) {
		var f PostFunc[In, Out] = func(ctx context.Context, in In, cr b.CommonRequest) Out {
			if _, ok := cr.QueryParams["param1"]; !ok {
				t.Fatalf("query param (param1) does not exist")
			}

			if _, ok := cr.QueryParams["param2"]; !ok {
				t.Fatalf("query param (param2) does not exist")
			}

			if in.Name != "tester" {
				t.Fatalf("name must be tester, got %s", in.Name)
			}

			if in.Age != 23 {
				t.Fatalf("age must be 23, got %d", in.Age)
			}

			var a Out

			a.Name = "tester"
			a.Age = 23
			return a
		}

		handler := POST(f, With(QueryParams("param1", "param2")))
		input := `{"name":"tester","age":23}`

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(input))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := methodServer.NewContext(req, rec)

		NoError(t, handler(c))
		Equal(t, 200, rec.Code)
	})
}

func TestPOSTErrorParseInput(t *testing.T) {
	t.Parallel()

	type In struct {
		b.CommonRequest
		Name string   `json:"name,omitempty"`
		Age  chan int `json:"Age,omitempty"`
	}

	type Out struct {
		b.CommonResponse
		Name string `json:"name,omitempty"`
		Age  int    `json:"age,omitempty"`
	}

	var f PostFunc[In, Out] = func(ctx context.Context, in In, cr b.CommonRequest) Out {
		var a Out

		a.Name = "tester"
		a.Age = 23
		return a
	}

	handler := POST(f, Default)
	input := `{"name":"tester","age":23}`

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(input))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := methodServer.NewContext(req, rec)

	NoError(t, handler(c))
	Equal(t, 400, rec.Code)
}
