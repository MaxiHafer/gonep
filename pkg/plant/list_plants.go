package plant

type ListPlantsResponse struct {
	Msg        string `json:"msg,omitempty"`
	StatusCode int    `json:"status,omitempty"`
	Data       struct {
		Plants []Plant `json:"list"`
		Page   int     `json:"page"`
		Total  int     `json:"totalpage"`
	} `json:"data"`
}
