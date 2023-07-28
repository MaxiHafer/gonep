package gonep

import (
	"context"
	"encoding/json"
	"github.com/maxihafer/gonep/internal"
	"net/http"
	"net/url"
	"strings"
)

const (
	NepViewerBasePath = "pv_monitor/appservice"
	LoginRequestPath  = "login"
)

var (
	defaultBaseUrl = &url.URL{
		Scheme: "http",
		Host:   "nep.nepviewer.com",
		Path:   "/pv_monitor/appservice",
	}

	defaultUsername = "anonymous"
)

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		BaseURL:    defaultBaseUrl,
		httpClient: http.DefaultClient,
		username:   defaultUsername,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

type Client struct {
	username string
	password string
	BaseURL  *url.URL

	httpClient *http.Client
	token      *string
}

func (c *Client) newRequest(ctx context.Context, method, path string, body *url.Values) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	if body == nil {
		body = &url.Values{}
	}

	if mustAuthenticate := c.token == nil; mustAuthenticate {
		if err := c.authenticate(); err != nil {
			return nil, err
		}
	}

	body.Set("token", *c.token)

	req, err := http.NewRequestWithContext(ctx, method, u.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return http.NewRequestWithContext(ctx, method, u.String(), strings.NewReader(body.Encode()))
}

func (c *Client) ListPVPlants(ctx context.Context) ([]PVPlant, error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/pvlist", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	plantResp := ListPlantsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&plantResp); err != nil {
		return nil, err
	}

	return plantResp.Plants, nil
}

func (c *Client) authenticate() error {
	rel := &url.URL{Path: "login"}
	u := c.BaseURL.ResolveReference(rel)

	body := &url.Values{}
	body.Set("email", c.username)
	body.Set("password", c.password)

	loginReq, err := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}
	loginReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(loginReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	loginResp := internal.LoginResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return err
	}

	if loginResp.Status != 1 {
		return UnsuccessfulLoginError{
			loginResp.Msg,
		}
	}

	c.token = &loginResp.Data.Token

	return nil
}
