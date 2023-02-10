package presentation

import (
	"context"
	"github.com/labstack/echo/v4"
	b "github.com/samuelsih/echo-structr/business"
	"net/http"
)

type GetFunc[out b.CommonOutput] func(ctx context.Context, cr b.CommonRequest) out
type PostFunc[inType any, out b.CommonOutput] func(ctx context.Context, in inType, cr b.CommonRequest) out

func GET[outType b.CommonOutput](businessFunc GetFunc[outType], opts Opts) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cr b.CommonRequest
		if len(opts.URLParams) > 0 {
			cr.URLParams = make(map[string]string, len(opts.URLParams))

			for _, param := range opts.URLParams {
				cr.URLParams[param] = c.Param(param)
			}
		}

		if len(opts.QueryParams) > 0 {
			cr.QueryParams = make(map[string]string, len(opts.QueryParams))

			for _, param := range opts.QueryParams {
				cr.QueryParams[param] = c.QueryParam(param)
			}
		}

		cookie, err := c.Cookie(cookieName)
		if err == http.ErrNoCookie {
			cr.Token = ""
		} else {
			cr.Token = cookie.Value
		}

		out := businessFunc(c.Request().Context(), cr)

		return RenderOutput(c, out, out.CommonResp())
	}
}

func POST[inType any, outType b.CommonOutput](businessFunc PostFunc[inType, outType], opts Opts) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			in inType
			cr b.CommonRequest
		)

		if err := ParseInput(c, &in, &cr); err != nil {
			resp := b.CommonResponse{
				StatusCode: 400,
				Msg:        "Bad Request",
			}

			return c.JSON(400, resp)
		}

		if len(opts.URLParams) > 0 {
			cr.URLParams = make(map[string]string, len(opts.URLParams))

			for _, param := range opts.URLParams {
				cr.URLParams[param] = c.Param(param)
			}
		}

		if len(opts.QueryParams) > 0 {
			cr.QueryParams = make(map[string]string, len(opts.QueryParams))

			for _, param := range opts.QueryParams {
				cr.QueryParams[param] = c.QueryParam(param)
			}
		}

		out := businessFunc(c.Request().Context(), in, cr)

		return RenderOutput(c, out, out.CommonResp())
	}
}
