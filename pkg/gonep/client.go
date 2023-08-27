package gonep

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/maxihafer/gonep/pkg/authentication"
	"github.com/maxihafer/gonep/pkg/gateway"
	"github.com/maxihafer/gonep/pkg/plant"
)

const (
	defaultScheme      = "http"
	defaultHost        = "nep.nepviewer.com"
	defaultServicePath = "pv_monitor/appservice"
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
