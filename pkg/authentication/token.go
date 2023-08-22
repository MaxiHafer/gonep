package authentication

import (
	"github.com/go-resty/resty/v2"
	"net/url"
)

var _ Provider = (*TokenProvider)(nil)

type TokenProvider struct {
	Token string
}

func (p *TokenProvider) Middleware() resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		body := url.Values{}
		body.Set("token", p.Token)

		request.SetFormDataFromValues(body)

		return nil
	}
}
