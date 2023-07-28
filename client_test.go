package gonep

import (
	"context"
	"github.com/stretchr/testify/suite"
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
	s.config = &Config{}
	//Tests require real credentials loaded into env
	s.Require().NoError(s.config.FromEnv())
}

func (s *ClientTestSuite) SetupTest() {
	s.client = NewClient(
		WithUserPassword(s.config.User, s.config.Password),
	)
}

func (s *ClientTestSuite) TestListPVPlants() {
	ctx := context.Background()

	_, err := s.client.ListPVPlants(ctx)
	s.Require().NoError(err)
}

func (s *ClientTestSuite) TestAuthenticate() {
	s.Require().NoError(s.client.authenticate())
}
