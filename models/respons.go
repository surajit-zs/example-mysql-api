package models

type Response struct {
	Cat        interface{} `json:"cat"`
	Message    string      `json:"massage"`
	StatusCode int         `json:"status_code"`
}
