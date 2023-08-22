package gonep

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/maxihafer/gonep/pkg/authentication"
	"github.com/maxihafer/gonep/pkg/plant"
)

const (
	defaultScheme      = "http"
	defaultHost        = "nep.nepviewer.com"
	defaultServicePath = "/pv_monitor/appservice"
)

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{}

	c.Client = resty.New()
	c.Client.SetScheme(defaultScheme)
	c.Client.SetBaseURL(fmt.Sprintf("%s://%s/%s", defaultScheme, defaultHost, defaultServicePath))

	for _, opt := range opts {
		opt(c)
	}

	if c.authenticationProvider == nil {
		return nil, errors.New("missing authentication provider must be set")
	}

	c.Client.OnBeforeRequest(c.authenticationProvider.Middleware())

	c.plantService = plant.NewService(c.Client)

	return c, nil
}

type Client struct {
	*resty.Client
	authenticationProvider authentication.Provider

	plantService plant.Service
}

func (c *Client) Plant() plant.Service {
	return c.plantService
}
