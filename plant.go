package gonep

import (
	"strconv"
	"time"
)

const (
	timeFormat = "2006-01-02 03:04"
)

type GatewayStatus struct {
	id        string
	Today     int
	Total     int
	CO2       int
	Current   int
	Status    string
	UpdatedAt time.Time
}

type PlantStatus struct {
	Current  int
	Today    int
	Year     int
	Total    int
	CO2      int
	Gateways []*GatewayStatus
}

type GetPlantStatusResponse struct {
	Today           int              `json:"today"`      //watt-hours
	Year            int              `json:"total_year"` //watt-hours
	Total           int              `json:"total"`      //watt-hours
	CO2             int              `json:"co2"`        //kilograms of co2 saved
	Current         int              `json:"now"`        //current energy production in watts
	Gateways        GatewayStatusMap `json:"gateway"`
	CircuitBoard    []string         `json:"circuit_board"`
	IsShow          int              `json:"is_show"`
	TemperatureUnit string           `json:"Temperature_Unit"`
	UpdatedAt       string           `json:"LastUpdate"`
}

func (r *GetPlantStatusResponse) PlantStatus() (*PlantStatus, error) {
	status := &PlantStatus{
		Current:  r.Current,
		Today:    r.Today,
		Year:     r.Year,
		Total:    r.Total,
		CO2:      r.CO2,
		Gateways: []*GatewayStatus{},
	}

	for id, v := range r.Gateways {

		ts, err := strconv.Atoi(v.UpdatedAt)
		if err != nil {
			return nil, err
		}

		status.Gateways = append(status.Gateways, &GatewayStatus{
			id:        id,
			Today:     v.Today,
			Total:     v.Total,
			CO2:       v.CO2,
			Current:   v.Current,
			Status:    v.Status,
			UpdatedAt: time.Unix(int64(ts), 0),
		})
	}

	return status, nil
}

type GatewayStatusMap map[string]gatewayStatus

type gatewayStatus struct {
	Today            int    `json:"today"`
	Total            int    `json:"total"`
	CO2              int    `json:"co2"`
	Current          int    `json:"now"`
	Status           string `json:"status"`
	RequestTimestamp int    `json:"NowD"`
	UpdatedAt        string `json:"LastUpdateTime"`
}

type ListPlantsResponse struct {
	Msg    string `json:"msg,omitempty"`
	Status int    `json:"status,omitempty"`
	Data   struct {
		Plants []Plant `json:"list"`
		Page   int     `json:"page"`
		Total  int     `json:"totalpage"`
	} `json:"data"`
}

type Plant struct {
	Sid               string   `json:"Sid,omitempty"`
	UserEmail         string   `json:"User_Email,omitempty"`
	SiteName          string   `json:"Site_Name,omitempty"`
	InstallationEmail string   `json:"Install_Email,omitempty"`
	Country           string   `json:"Country,omitempty"`
	State             string   `json:"State,omitempty"`
	City              string   `json:"City,omitempty"`
	Street            string   `json:"Street,omitempty"`
	Latitude          string   `json:"Latitude,omitempty"`
	Longitude         string   `json:"Longitude,omitempty"`
	ImageRef          string   `json:"Logo,omitempty"`
	GroupStrSN        string   `json:"Group_strSN,omitempty"`
	TemperatureUnit   string   `json:"Temperature_Unit,omitempty"`
	Gateways          []string `json:"Gateways,omitempty"`
}
