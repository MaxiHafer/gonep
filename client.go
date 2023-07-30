package gonep

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/maxihafer/gonep/internal"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
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
		Path:   "/pv_monitor/appservice/",
	}

	defaultUsername = "anonymous"
)

func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		BaseURL:    defaultBaseUrl,
		httpClient: http.DefaultClient,
		username:   defaultUsername,
	}

	var err error
	c.httpClient.Jar, err = cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
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

	return req, nil
}

func (c *Client) GetPlantStatus(ctx context.Context, sid string) (*PlantStatus, error) {
	req, err := c.newRequest(ctx, http.MethodPost, fmt.Sprintf("status/%s", sid), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	statusResp := GetPlantStatusResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&statusResp); err != nil {
		return nil, err
	}

	return statusResp.PlantStatus()
}

func (c *Client) ListPlants(ctx context.Context) ([]Plant, error) {
	req, err := c.newRequest(ctx, http.MethodPost, "pvlist", nil)
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

	if plantResp.Status != 1 {
		return nil, errors.New(fmt.Sprintf("received code: %v, message: %s", plantResp.Status, plantResp.Msg))
	}

	return plantResp.Data.Plants, nil
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
