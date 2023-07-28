package gonep

import (
	"net/http"
	"net/url"
)

type ClientOption func(*Client)

func WithBaseURL(url *url.URL) ClientOption {
	return func(client *Client) {
		client.BaseURL = url
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func WithUserPassword(username, password string) ClientOption {
	return func(client *Client) {
		client.username = username
		client.password = password
	}
}
