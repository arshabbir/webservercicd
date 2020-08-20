package utils

type APIerror struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}
