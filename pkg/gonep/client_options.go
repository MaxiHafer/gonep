package gonep

import (
	"github.com/maxihafer/gonep/pkg/authentication"
	"net/url"
)

type ClientOption func(*Client)

func WithBaseURL(url *url.URL) ClientOption {
	return func(client *Client) {
		client.Client.SetBaseURL(url.String())
	}
}

func WithUserPassword(username, password string) ClientOption {
	return func(client *Client) {
		client.authenticationProvider = &authentication.BasicProvider{
			Username: username,
			Password: password,
		}
	}
}

func WithToken(token string) ClientOption {
	return func(client *Client) {
		client.authenticationProvider = &authentication.TokenProvider{
			Token: token,
		}
	}
}
