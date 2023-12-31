package plant

import (
	"encoding/json"
	"github.com/maxihafer/gonep/pkg/gateway"
)

const (
	timeFormat = "2006-01-02 03:04"
)

var _ json.Unmarshaler = (*Status)(nil)

type Status struct {
	CurrentWatts        int
	TodayWattHours      int
	YearWattHours       int
	TotalWattHours      int
	KilogramsOfCO2Saved int
	Gateways            []*gateway.Status
}

func (s *Status) UnmarshalJSON(data []byte) error {
	type rawGatewayStatus struct {
		Today  int    `json:"today"`
		Total  int    `json:"total"`
		CO2    int    `json:"co2"`
		Now    int    `json:"now"`
		Status string `json:"status"`
	}

	type rawResponse struct {
		Today           int                         `json:"today"`
		Year            int                         `json:"total_year"`
		Total           int                         `json:"total"`
		CO2             int                         `json:"co2"`
		Now             int                         `json:"now"`
		Gateway         map[string]rawGatewayStatus `json:"gateway"`
		CircuitBoard    []string                    `json:"circuit_board"`
		IsShow          int                         `json:"is_show"`
		TemperatureUnit string                      `json:"Temperature_Unit"`
		LastUpdate      string                      `json:"LastUpdate"`
	}

	raw := new(rawResponse)
	if err := json.Unmarshal(data, raw); err != nil {
		return err
	}

	s.CurrentWatts = raw.Now
	s.TodayWattHours = raw.Today
	s.YearWattHours = raw.Year
	s.TotalWattHours = raw.Total
	s.KilogramsOfCO2Saved = raw.CO2

	for id, sts := range raw.Gateway {
		s.Gateways = append(s.Gateways, &gateway.Status{
			Id:                  id,
			TodayWattHours:      sts.Today,
			TotalWattHours:      sts.Total,
			KilogramsOfCO2Saved: sts.CO2,
			CurrentWatts:        sts.Now,
			Status:              sts.Status,
		})
	}

	return nil
}
