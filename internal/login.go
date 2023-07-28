package internal

type LoginResponse struct {
	Msg    string            `json:"msg"`
	Status int               `json:"status"`
	Data   LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	Token    string `json:"token"`
	Language string `json:"lang"`
	IsOpen   int    `json:"isopen"`
}
