package authentication

import (
	"github.com/go-resty/resty/v2"
)

type Provider interface {
	Middleware() resty.RequestMiddleware
}
