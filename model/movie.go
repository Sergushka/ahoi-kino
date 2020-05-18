package model

import "time"

type Movie struct {
	Id                  int              `json:"id"`
	Adult               bool             `json:"adult"`
	BackdropPath        string           `json:"backdrop_path"`
	BelongsToCollection Collection       `json:"belongs_to_collection"`
	Budget              int              `json:"budget"`
	CreatedAt           time.Duration    `json:"created_at"`
	DirectLink          string           `json:"direct_link"`
	Genres              []Genre          `json:"genres"`
	Homepage            string           `json:"homepage"`
	ImdbId              string           `json:"imdb_id"`
	Name                string           `json:"name"`
	OriginalLanguage    string           `json:"original_language"`
	OriginalTitle       string           `json:"original_title"`
	Overview            string           `json:"overview"`
	Popularity          float64          `json:"popularity"`
	PosterPath          string           `json:"poster_path"`
	ProductionCompanies []Company        `json:"production_companies"`
	ProductionCountries []Country        `json:"production_countries"`
	ReleaseDate         string           `json:"release_date"`
	Revenue             int              `json:"revenue"`
	Runtime             int              `json:"runtime"`
	Screens             []ScreenShot     `json:"screens"`
	SpokenLanguages     []SpokenLanguage `json:"spoken_languages"`
	Status              string           `json:"status"`
	Tagline             string           `json:"tagline"`
	Title               string           `json:"title"`
	TmdbId              int              `json:"tmdb_id"`
	UserUid             string           `json:"user_uid"`
	Video               bool             `json:"video"`
	VoteAverage         float64          `json:"vote_average"`
	VoteCount           int              `json:"vote_count"`
}
