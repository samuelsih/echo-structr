package presentation

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/samuelsih/echo-structr/business"
	"net/http"
)

const (
	cookieName    string = "token"
	cookieExpired int    = 60 * 60 * 24 // 1 day
)

func ParseInput[T any](ctx echo.Context, in T, cr *business.CommonRequest) error {
	err := ctx.Bind(in)
	if err != nil {
		return err
	}

	cookie, err := ctx.Cookie(cookieName)
	if errors.Is(err, http.ErrNoCookie) {
		cr.Token = ""
		return nil
	}

	cr.Token = cookie.Value
	return nil
}

func RenderOutput[T any](ctx echo.Context, out T, cr business.CommonResponse) error {
	if cr.StatusCode == 0 {
		cr.StatusCode = 200
	}

	if cr.SetAuthToken != "" {
		cookie := &http.Cookie{
			Name:     cookieName,
			Value:    cr.SetAuthToken,
			MaxAge:   cookieExpired,
			HttpOnly: true,
			SameSite: 0,
		}

		ctx.SetCookie(cookie)
	}

	return ctx.JSON(cr.StatusCode, out)
}
