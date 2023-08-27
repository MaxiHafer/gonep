package gateway

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
	"go.nhat.io/httpmock"
	"testing"
	"time"
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
	client.SetBaseURL(fmt.Sprintf("%s/pv_monitor/appservice", s.server.URL()))
	client.SetScheme("http")

	s.service = &service{
		client: client,
	}
}

var (
	todayResponseData = []byte(`[
		[1693118280000,0],
		[1693189980000,29],
		[1693170001000,null]
	]`)
)

func (s *ServiceTestSuite) TestToday() {
	s.server.ExpectPost("/pv_monitor/appservice/detail/1337").Return(todayResponseData)

	metrics, err := s.service.Today(context.Background(), "1337")
	s.Require().NoError(err)

	s.Require().Contains(metrics, &Metric{
		Timestamp: time.UnixMilli(1693118280000),
		Watts:     0,
	})

	s.Require().Contains(metrics, &Metric{
		Timestamp: time.UnixMilli(1693189980000),
		Watts:     29,
	})

	s.Require().Contains(metrics, &Metric{
		Timestamp: time.UnixMilli(1693170001000),
		Watts:     0,
	})
}
