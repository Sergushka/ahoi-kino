package model

type MoviesRequest struct {
	MovieCount int `json:"movieCount"`
	Page       int `json:"page"`
}
