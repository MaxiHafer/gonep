package test

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/k0kubun/pp/v3"
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
	s.Require().NoError(godotenv.Load("../.env"))

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

	pp.Println("Received Plants:")
	for _, plant := range plants {
		pp.Println(plant)
	}

	status, err := s.client.Plants().Status(ctx, plants[0].Sid)
	s.Require().NoError(err)
	s.Require().NotNil(status)

	gw := status.Gateways[0].Id

	pp.Printf("Received Status for Plant %s:\n", plants[0].Sid)
	pp.Println(status)

	today, err := s.client.Gateways().Today(ctx, gw)
	s.Require().NoError(err)
	s.Require().NotEmpty(today)

	pp.Printf("Received today metrics for gateway %s", gw)
	pp.Println(today)

	week, err := s.client.Gateways().Week(ctx, gw)
	s.Require().NoError(err)
	s.Require().NotEmpty(week)

	pp.Printf("Received week metrics for gateway %s", gw)
	pp.Println(week)

	month, err := s.client.Gateways().Month(ctx, gw)
	s.Require().NoError(err)
	s.Require().NotEmpty(month)

	pp.Printf("Received month metrics for gateway %s", gw)
	pp.Println(month)

	year, err := s.client.Gateways().Year(ctx, gw)
	s.Require().NoError(err)
	s.Require().NotEmpty(year)

	pp.Printf("Received year metrics for gateway %s", gw)
	pp.Println(year)
}
