package model

type Response struct {
	Err    string      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}
