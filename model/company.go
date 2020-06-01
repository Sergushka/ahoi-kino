package model

type Company struct {
	Id             int    `json:"id"`
	Logo_Path      string `json:"logo_path"`
	Name           string `json:"name"`
	Origin_Country string `json:"origin_country"`
}
