package gateway

type Status struct {
	Today            int    `json:"today"`
	Total            int    `json:"total"`
	CO2              int    `json:"co2"`
	Current          int    `json:"now"`
	Status           string `json:"status"`
	RequestTimestamp int    `json:"NowD"`
	UpdatedAt        string `json:"LastUpdateTime"`
}
