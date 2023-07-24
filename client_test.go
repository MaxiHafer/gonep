package gonep

import (
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

type ClientTestSuite struct {
	suite.Suite

	client *Client
	config *Config
}

func (s *ClientTestSuite) SetupSuite() {
	//This requires real credentials loaded into env
	s.Require().NoError(s.config.FromEnv())
}

func (s *ClientTestSuite) SetupTest() {
	s.client = NewClient(s.config)

	s.Require().NoError(s.client.Init())
}

func (s *ClientTestSuite) TestClientInit() {
	cookies := s.client.Jar.Cookies(&url.URL{
		Scheme: s.client.config.Scheme,
		Host:   s.client.config.BaseURL,
	})

	s.Require().Equal(cookies[0].Name, "PHPSESSID")
}

func (s *ClientTestSuite) TestGetCaptcha() {
	captcha, err := s.client.getCaptcha()
	s.Require().NoError(err)
	s.Require().Greater(*captcha, 0)
	s.Require().Less(*captcha, 10000)
}
