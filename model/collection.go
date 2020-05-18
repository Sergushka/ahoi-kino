package model

type Collection struct {
	BackdropPath string `json:"backdrop_path"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	PosterPath   string `json:"poster_path"`
}
