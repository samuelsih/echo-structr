package business

// type check interface at compile time
var _ CommonOutput = (*CommonResponse)(nil)

type CommonRequest struct {
	Token       string            `json:"token,omitempty"`
	URLParams   map[string]string `json:"-"`
	QueryParams map[string]string `json:"-"`
}

type CommonResponse struct {
	SetAuthToken string `json:"-"`
	StatusCode   int    `json:"status_code,omitempty"`
	Msg          string `json:"msg,omitempty"`
}

func (cr *CommonResponse) Set(code int, msg string) {
	cr.StatusCode = code
	cr.Msg = msg
}

func (cr CommonResponse) CommonResp() CommonResponse {
	return cr
}

type CommonOutput interface {
	CommonResp() CommonResponse
}
