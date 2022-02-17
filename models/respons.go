package models

type Response struct {
	Cat        interface{} `json:"cat"`
	Massage    string      `json:"massage"`
	StatusCode int         `json:"status_code"`
}
