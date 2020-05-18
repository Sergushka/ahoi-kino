package model

type Country struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}
