package gonep

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
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
	var err error
	s.client, err = NewClient(
		WithUserPassword(s.config.User, s.config.Password),
	)
	s.Require().NoError(err)
}

func (s *ClientTestSuite) TestListPVPlants() {
	ctx := context.Background()

	plants, err := s.client.Plants().List(ctx)
	s.Require().NoError(err)
	logrus.WithField("plants", fmt.Sprintf("%+v", plants)).Info("Got plants")
}

func (s *ClientTestSuite) TestGetPlantStatus() {
	ctx := context.Background()

	plants, err := s.client.Plants().List(ctx)
	s.Require().NoError(err)

	status, err := s.client.Plants().Status(ctx, plants[0].Sid)
	s.Require().NoError(err)
	s.Require().NotNil(status)

	fmt.Printf("%#v", status)
}
