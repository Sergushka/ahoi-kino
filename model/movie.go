package model

type Movie struct {
	Id                    string           `json:"id,omitempty"`
	Adult                 bool             `json:"adult,omitempty"`
	Backdrop_Path         string           `json:"backdrop_path,omitempty"`
	Belongs_To_Collection *Collection      `json:"belongs_to_collection,omitempty"`
	Budget                int              `json:"budget,omitempty"`
	CreatedAt             int              `json:"created_at,omitempty"`
	DirectLink            string           `json:"direct_link,omitempty"`
	Genres                []Genre          `json:"genres,omitempty"`
	Homepage              string           `json:"homepage,omitempty"`
	Imdb_Id               string           `json:"imdb_id,omitempty"`
	Name                  string           `json:"name,omitempty"`
	Original_Language     string           `json:"original_language,omitempty"`
	Original_Title        string           `json:"original_title,omitempty"`
	Overview              string           `json:"overview,omitempty"`
	Popularity            float64          `json:"popularity,omitempty"`
	Poster_Path           string           `json:"poster_path,omitempty"`
	Production_Companies  []Company        `json:"production_companies,omitempty"`
	Production_Countries  []Country        `json:"production_countries,omitempty"`
	Release_Date          string           `json:"release_date,omitempty"`
	Revenue               int              `json:"revenue,omitempty"`
	Runtime               int              `json:"runtime,omitempty"`
	Screens               []ScreenShot     `json:"screens,omitempty"`
	Spoken_Languages      []SpokenLanguage `json:"spoken_languages,omitempty"`
	Status                string           `json:"status,omitempty"`
	Tagline               string           `json:"tagline,omitempty"`
	Title                 string           `json:"title,omitempty"`
	Tmdb_Id               int              `json:"tmdb_id,omitempty"`
	UserUid               string           `json:"user_uid,omitempty"`
	Video                 bool             `json:"video,omitempty"`
	Vote_Average          float64          `json:"vote_average,omitempty"`
	Vote_Count            int              `json:"vote_count,omitempty"`
}
