package test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/maxihafer/gonep/pkg/client"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

type IntegrationTestSuite struct {
	suite.Suite

	client *client.Client
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.Require().NoError(godotenv.Load())

	host, ok := os.LookupEnv("NEPVIEWER_URL")
	s.Require().True(ok)

	email, ok := os.LookupEnv("NEPVIEWER_EMAIL")
	s.Require().True(ok)

	password, ok := os.LookupEnv("NEPVIEWER_PASSWORD")
	s.Require().True(ok)

	var err error
	s.client, err = client.NewClient(
		client.WithBaseURL(host),
		client.WithEmailPassword(email, password),
	)

	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) Test() {
	ctx := context.Background()

	plants, err := s.client.Plants().List(ctx)
	s.Require().NoError(err)
	s.Require().NotEmpty(plants)

	fmt.Println(plants)

	status, err := s.client.Plants().Status(ctx, plants[0].Sid)
	s.Require().NoError(err)
	s.Require().NotNil(status)

	gw := status.Gateways[0].Id

	fmt.Println(status)

	today, err := s.client.Gateways().Today(ctx, gw)
	s.Require().NoError(err)
	s.Require().NotEmpty(today)

	fmt.Println(today)

	week, err := s.client.Gateways().Week(ctx, gw)
	s.Require().NoError(err)
	s.Require().NotEmpty(week)

	fmt.Println(week)

	month, err := s.client.Gateways().Month(ctx, gw)
	s.Require().NoError(err)
	s.Require().NotEmpty(month)

	fmt.Println(month)

	year, err := s.client.Gateways().Year(ctx, gw)
	s.Require().NoError(err)
	s.Require().NotEmpty(year)

	fmt.Println(year)
}
