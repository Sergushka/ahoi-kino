package model

type Response struct {
	Err    error       `json:"error"`
	Result interface{} `json:"result"`
}
