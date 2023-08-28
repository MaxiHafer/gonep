package authentication

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	"go.nhat.io/httpmock"
	"net/url"
	"testing"
)

func TestTokenProviderTestSuite(t *testing.T) {
	suite.Run(t, new(TokenProviderTestSuite))
}

type TokenProviderTestSuite struct {
	suite.Suite

	server *httpmock.Server
}

func (s *TokenProviderTestSuite) SetupSuite() {
	s.server = httpmock.NewServer()
}

func (s *TokenProviderTestSuite) TestMiddleware() {
	provider := TokenProvider{
		Token: "1234",
	}

	tokenBody := url.Values{}
	tokenBody.Set("token", "1234")

	s.server.ExpectPost("/").WithBody(tokenBody.Encode()).Once()

	client := resty.New()
	client.SetBaseURL(s.server.URL())
	client.SetScheme("http")
	client.OnBeforeRequest(provider.Middleware())

	_, err := client.R().Post("/")
	s.Require().NoError(err)
}
