package client

import (
	"errors"
	"github.com/go-resty/resty/v2"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/maxihafer/gonep/pkg/authentication"
	"github.com/maxihafer/gonep/pkg/gateway"
	"github.com/maxihafer/gonep/pkg/plant"
)

const (
	defaultHost = "https://nep.nepviewer.com"
)

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{}

	c.Client = resty.New()
	c.SetBaseURL(defaultHost)

	for _, opt := range opts {
		opt(c)
	}

	if c.authenticationProvider == nil {
		return nil, errors.New("one authentication provider must be set")
	}

	c.Client.OnBeforeRequest(c.authenticationProvider.Middleware())

	c.plantService = plant.NewService(c.Client)
	c.gatewayService = gateway.NewService(c.Client)

	return c, nil
}

type Client struct {
	*resty.Client
	authenticationProvider authentication.Provider

	plantService   plant.Service
	gatewayService gateway.Service
}

func (c *Client) Plants() plant.Service {
	return c.plantService
}

func (c *Client) Gateways() gateway.Service {
	return c.gatewayService
}
