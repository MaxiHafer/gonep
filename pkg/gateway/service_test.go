package gateway

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	"go.nhat.io/httpmock"
	"testing"
)

var (
	timestampResponseBody = []byte(`[
	  [
		1693161960000,
		29
	  ],
	  [
		1693162260000,
		23
	  ]
	]`)

	monthResponseBody = []byte(`[
	  [
		2023.01,
		29
	  ],
	  [
		2023.02,
		23
	  ]
	]`)

	dayResponseBody = []byte(`[
	  [
		"01\/01",
		0.3
	  ],
	  [
		"01\/02",
		1.8
	  ]
	]`)

	yearResponseBody = []byte(`[
	  [
		2022.01,
		29
	  ],
	  [
		2023.01,
		23
	  ]
	]`)
)

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite

	server  *httpmock.Server
	service *service
}

func (s *ServiceTestSuite) SetupSuite() {
	s.server = httpmock.New()(s.T())

	client := resty.New()
	client.SetBaseURL(s.server.URL())
	client.SetScheme("http")

	s.service = &service{
		client: client,
	}
}

func (s *ServiceTestSuite) TestToday() {
	s.server.ExpectPost("/pv_monitor/appservice/detail/1337/0").Return(timestampResponseBody)

	metrics, err := s.service.Today(context.Background(), "1337")
	s.Require().NoError(err)

	s.Require().NotEmpty(metrics)
}

func (s *ServiceTestSuite) TestWeek() {
	s.server.ExpectPost("/pv_monitor/appservice/week/1337/0").Return(dayResponseBody)

	metrics, err := s.service.Week(context.Background(), "1337")
	s.Require().NoError(err)

	s.Require().NotEmpty(metrics)
}

func (s *ServiceTestSuite) TestMonth() {
	s.server.ExpectPost("/pv_monitor/appservice/month/1337/0").Return(monthResponseBody)

	metrics, err := s.service.Month(context.Background(), "1337")
	s.Require().NoError(err)

	s.Require().NotEmpty(metrics)
}

func (s *ServiceTestSuite) TestYear() {
	s.server.ExpectPost("/pv_monitor/appservice/year/1337/0").Return(yearResponseBody)

	metrics, err := s.service.Year(context.Background(), "1337")
	s.Require().NoError(err)

	s.Require().NotEmpty(metrics)
}
