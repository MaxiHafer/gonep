package authentication

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/maxihafer/gonep/pkg/authentication/internal"
	"net/url"
)

var _ Provider = (*BasicProvider)(nil)

type BasicProvider struct {
	Email    string
	Password string

	token *string
}

func (p *BasicProvider) Middleware() resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		if mustAuthenticate := p.token == nil; mustAuthenticate {
			token, err := p.fetchToken(request.Context(), client)

			if err != nil {
				return err
			}

			p.token = token
		}

		body := url.Values{}
		body.Set("token", *p.token)

		request.SetFormDataFromValues(body)

		return nil
	}
}

func (p *BasicProvider) fetchToken(ctx context.Context, client *resty.Client) (*string, error) {
	body := url.Values{}
	body.Set("email", p.Email)
	body.Set("password", p.Password)

	// use copy because we cannot disable the authentication middleware
	authClient := resty.NewWithClient(client.GetClient())
	authClient.SetBaseURL(client.BaseURL)
	authClient.SetScheme("http")

	resp, err := authClient.R().
		SetContext(ctx).
		SetFormDataFromValues(body).
		SetHeader("Cache-Control", "no-cache").
		Post("/pv_monitor/appservice/login")

	if err != nil {
		return nil, err
	}

	loginResponse := new(internal.LoginResponse)
	if err := json.Unmarshal(resp.Body(), loginResponse); err != nil {
		return nil, err
	}

	if loginResponse.Status != 1 {
		return nil, fmt.Errorf("authentication/basic: error status code '%v', message: '%s'", loginResponse.Status, loginResponse.Msg)
	}

	return &loginResponse.Data.Token, nil
}
