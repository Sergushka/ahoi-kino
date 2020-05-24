package model

type TMDBMovieResponse struct {
	PosterPath       string  `json:"poster_path,omitempty"`
	Popularity       float64 `json:"popularity,omitempty"`
	VoteCount        int     `json:"vote_count,omitempty"`
	Video            bool    `json:"video,omitempty"`
	MediaType        string  `json:"media_type,omitempty"`
	Id               int     `json:"id,omitempty"`
	Adult            bool    `json:"adult,omitempty"`
	BackdropPath     string  `json:"backdrop_path,omitempty"`
	OriginalLanguage string  `json:"original_language,omitempty"`
	OriginalTitle    string  `json:"original_title,omitempty"`
	GenreIds         []int   `json:"genre_ids,omitempty"`
	Title            string  `json:"title,omitempty"`
	VoteAverage      float64 `json:"vote_average,omitempty"`
	Overview         string  `json:"overview,omitempty"`
	ReleaseDate      string  `json:"release_date,omitempty"`
}
