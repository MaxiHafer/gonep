package gonep

import (
	"bufio"
	"encoding/binary"
	"github.com/maxihafer/gonep/pkg/pointer"
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func NewInitializedClientFromEnv() (*Client, error) {
	client, err := NewClientFromEnv()
	if err != nil {
		return nil, err
	}

	if err = client.Init(); err != nil {
		return nil, err
	}

	return client, nil
}

func NewInitializedClient(config *Config) (*Client, error) {
	client := NewClient(config)

	if err := client.Init(); err != nil {
		return nil, err
	}

	return client, nil
}

func NewClientFromEnv() (*Client, error) {
	config := &Config{}
	if err := config.FromEnv(); err != nil {
		return nil, err
	}

	return NewClient(config), nil
}

func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		Client: &http.Client{},
	}
}

type Client struct {
	config *Config
	*http.Client

	token *string
}

func (c *Client) Init() error {
	cookieJar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return err
	}
	c.Jar = cookieJar

	reqUrl := url.URL{
		Host:   c.config.BaseURL,
		Scheme: c.config.Scheme,
		Path:   "pv_manager/pv.php",
	}

	resp, err := c.Get(reqUrl.String())
	if err != nil {
		return err
	}

	return resp.Body.Close()
}

func (c *Client) getCaptcha() (*int, error) {
	reqUrl := &url.URL{
		Host:   c.config.BaseURL,
		Scheme: c.config.Scheme,
		Path:   "/pv_manager/captcha.php",
	}

	resp, err := c.Get(reqUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	captcha, err := binary.ReadVarint(bufio.NewReader(resp.Body))
	if err != nil {
		return nil, err
	}

	return pointer.Of(int(captcha)), nil
}

func (c *Client) login() error {
	captcha, err := c.getCaptcha()
	if err != nil {
		return err
	}

	reqUrl := &url.URL{
		Scheme: c.config.Scheme,
		Host:   c.config.BaseURL,
		Path:   "/pv_manager/login.php",
	}

	// TODO: Do auth magic rel: https://lanbugs.de/go-golang-co/go-http-client-post-form-data-x-www-form-urlencoded/
}

func (c *Client) do() (*http.Response, error) {
	if c.token == nil {
		//captcha, err := c.getCaptcha()
		if err := c.login(); err != nil {
			return nil, err
		}
	}
}
