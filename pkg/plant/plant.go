package plant

import (
	"encoding/json"
	"fmt"
)

var _ json.Unmarshaler = (*Plant)(nil)

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

type PlantList []*Plant

func (p *PlantList) UnmarshalJSON(data []byte) error {
	type rawResponse struct {
		Msg        string `json:"msg,omitempty"`
		StatusCode int    `json:"status,omitempty"`
		Data       struct {
			Plants []*Plant `json:"list"`
			Page   int      `json:"page"`
			Total  int      `json:"totalpage"`
		} `json:"data"`
	}

	raw := new(rawResponse)
	if err := json.Unmarshal(data, raw); err != nil {
		return err
	}

	for i := len(raw.Data.Plants) {

	}
}

func (p *Plant) String() string {
	return fmt.Sprintf("Sid:%s")
}
