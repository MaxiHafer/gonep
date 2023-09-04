package plant

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/maxihafer/gonep/pkg/gateway"
	"github.com/stretchr/testify/suite"
	"go.nhat.io/httpmock"
	"testing"
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

	s.service = &service{
		client: client,
	}
}

var (
	statusResponseData = []byte(`{
	  "today": 1748,
	  "total_status": 89065,
	  "total_year": 87317,
	  "total": 89065,
	  "co2": 96,
	  "now": 0,
	  "photovoltaic_panel": {
		"is_show": 1,
		"value": 0,
		"unit": "KWh",
		"direction": 0
	  },
	  "power_grid": {
		"is_show": 1,
		"value": 0,
		"unit": "KWh",
		"direction": 0
	  },
	  "home": {
		"is_show": 1,
		"value": 0,
		"unit": "KWh",
		"direction": 0
	  },
	  "battery": {
		"is_show": 1,
		"value": 0,
		"unit": "KWh",
		"direction": 0,
		"soc_value": "0%"
	  },
	  "LastUpdate": "2023-08-27 20:15",
	  "gateway": {
		"32c833a0": {
		  "today": 837,
		  "total": 43366,
		  "co2": 47,
		  "now": 0,
		  "photovoltaic_panel": {
			"is_show": 1,
			"value": 0,
			"unit": "KWh",
			"direction": 0
		  },
		  "power_grid": {
			"is_show": 1,
			"value": 0,
			"unit": "KWh",
			"direction": 0
		  },
		  "home": {
			"is_show": 1,
			"value": 0,
			"unit": "KWh",
			"direction": 0
		  },
		  "battery": {
			"is_show": 1,
			"value": 0,
			"unit": "KWh",
			"direction": 0,
			"soc_value": "0%"
		  },
		  "status": "0000",
		  "NowD": 1693165647,
		  "LastUpdateTime": "1693167344",
		  "LastUpdate": "2023-08-27 20:15",
		  "difference": -1697
		},
		"32c800c0": {
		  "today": 911,
		  "total": 45699,
		  "co2": 49,
		  "now": 0,
		  "photovoltaic_panel": {
			"is_show": 1,
			"value": 0,
			"unit": "KWh",
			"direction": 0
		  },
		  "power_grid": {
			"is_show": 1,
			"value": 0,
			"unit": "KWh",
			"direction": 0
		  },
		  "home": {
			"is_show": 1,
			"value": 0,
			"unit": "KWh",
			"direction": 0
		  },
		  "battery": {
			"is_show": 1,
			"value": 0,
			"unit": "KWh",
			"direction": 0,
			"soc_value": "0%"
		  },
		  "status": "0000",
		  "NowD": 1693165647,
		  "LastUpdateTime": "1693167352",
		  "LastUpdate": "2023-08-27 20:15",
		  "difference": -1705
		}
	  },
	  "circuit_board": [],
	  "is_show": 0,
	  "Temperature_Unit": "Celsius"
	}`)

	listResponseData = []byte(`{
	  "msg": "",
	  "status": 1,
	  "data": {
		"list": [
		  {
			"Sid": "TEST_20230827_xY47",
			"User_Email": "testuser@example.com",
			"Site_Name": "Test Home",
			"Install_Email": "installer@example.com",
			"Country": "US",
			"State": "California",
			"City": "Los Angeles",
			"Street": "123 Main St",
			"Latitude": "34.052235",
			"Longitude": "-118.243683",
			"Logo": "http://example.com/testlogo.jpg",
			"Group_strSN": "aBcDeF12,GhIjKl34",
			"Temperature_Unit": "Celsius",
			"Gateways": ["aBcDeF12", "GhIjKl34"]
		  }
		],
		"totalpage": 1,
		"page": 1
	  }
	}`)
)

func (s *ServiceTestSuite) TestToday() {
	s.server.ExpectPost("/pv_monitor/appservice/status/1337").Return(statusResponseData)

	status, err := s.service.Status(context.Background(), "1337")
	s.Require().NoError(err)

	s.Require().Equal(status.CurrentWatts, 0)
	s.Require().Equal(status.TodayWattHours, 1748)
	s.Require().Equal(status.TotalWattHours, 89065)
	s.Require().Equal(status.YearWattHours, 87317)
	s.Require().Equal(status.KilogramsOfCO2Saved, 96)

	s.Require().Contains(status.Gateways, &gateway.Status{
		Id:                  "32c800c0",
		CurrentWatts:        0,
		TodayWattHours:      911,
		TotalWattHours:      45699,
		KilogramsOfCO2Saved: 49,
		Status:              "0000",
	})
	s.Require().Contains(status.Gateways, &gateway.Status{
		Id:                  "32c833a0",
		CurrentWatts:        0,
		TodayWattHours:      837,
		TotalWattHours:      43366,
		KilogramsOfCO2Saved: 47,
		Status:              "0000",
	})
}

func (s *ServiceTestSuite) TestList() {
	s.server.ExpectPost("/pv_monitor/appservice/pvlist").Return(listResponseData)

	plants, err := s.service.List(context.Background())
	s.Require().NoError(err)

	s.Require().Contains(plants, &Plant{
		Sid:               "TEST_20230827_xY47",
		UserEmail:         "testuser@example.com",
		SiteName:          "Test Home",
		InstallationEmail: "installer@example.com",
		Country:           "US",
		State:             "California",
		City:              "Los Angeles",
		Street:            "123 Main St",
		Latitude:          "34.052235",
		Longitude:         "-118.243683",
		ImageRef:          "http://example.com/testlogo.jpg",
		GroupStrSN:        "aBcDeF12,GhIjKl34",
		TemperatureUnit:   "Celsius",
		Gateways: []string{
			"aBcDeF12",
			"GhIjKl34",
		},
	})
}
