package client

import (
	"github.com/maxihafer/gonep/pkg/authentication"
)

type Option func(*Client)

func WithBaseURL(host string) Option {
	return func(client *Client) {
		client.Client.SetBaseURL(host)
	}
}

func WithEmailPassword(email, password string) Option {
	return func(client *Client) {
		client.authenticationProvider = &authentication.BasicProvider{
			Email:    email,
			Password: password,
		}
	}
}

func WithToken(token string) Option {
	return func(client *Client) {
		client.authenticationProvider = &authentication.TokenProvider{
			Token: token,
		}
	}
}
