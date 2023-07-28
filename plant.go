package gonep

type ListPlantsResponse struct {
	Mg     string    `json:"msg,omitempty"`
	Status string    `json:"status,omitempty"`
	Plants []PVPlant `json:"plants,omitempty"`
}

type PVPlant struct {
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
