package authentication

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/maxihafer/gonep/pkg/authentication/internal"
	"github.com/maxihafer/gonep/pkg/pointer"
	"github.com/stretchr/testify/suite"
	"go.nhat.io/httpmock"
	"net/url"
	"testing"
)

func TestBasicProviderTestSuite(t *testing.T) {
	suite.Run(t, new(BasicProviderTestSuite))
}

type BasicProviderTestSuite struct {
	suite.Suite

	server *httpmock.Server
}

func (s *BasicProviderTestSuite) SetupSuite() {
	s.server = httpmock.NewServer()
}

func (s *BasicProviderTestSuite) TestMiddleware_TokenSet() {
	provider := BasicProvider{
		token: pointer.Of("1234"),
	}

	client := resty.New()
	client.OnBeforeRequest(provider.Middleware())

	body := url.Values{}
	body.Set("token", "1234")

	s.server.ExpectPost("/").WithBody(body.Encode())

	_, err := client.R().Post(fmt.Sprintf("%s/", s.server.URL()))
	s.Require().NoError(err)
}

func (s *BasicProviderTestSuite) TestMiddleware_TokenUnset() {
	provider := BasicProvider{
		Email:    "test",
		Password: "1234",
	}

	loginBody := url.Values{}
	loginBody.Set("email", "test")
	loginBody.Set("password", "1234")

	s.server.
		ExpectPost("/pv_monitor/appservice/login").
		WithBody(loginBody.Encode()).
		ReturnJSON(internal.LoginResponse{
			Msg:    "",
			Status: 1,
			Data: internal.LoginResponseData{
				Token: "1234",
			},
		}).Once()

	postBody := url.Values{}
	postBody.Set("token", "1234")

	s.server.
		ExpectPost("/pv_monitor/appservice/test").
		WithBody(postBody.Encode()).
		Once()

	client := resty.New()
	client.SetBaseURL(s.server.URL())
	client.OnBeforeRequest(provider.Middleware())

	resp, err := client.R().Post("/pv_monitor/appservice/test")
	s.Require().NoError(err)

	s.Require().NotNil(resp)

}
