package web

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Massage string      `json:"massage"`
	Data    interface{} `json:"data"`
}
