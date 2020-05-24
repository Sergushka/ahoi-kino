package graph

import "github.com/graphql-go/graphql"

var collectionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Collection",
		Description: "Collection movie belongs to",
		Fields: graphql.Fields{
			"backdrop_path": &graphql.Field{
				Type: graphql.String,
			},
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"poster_path": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var genreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Genre",
		Description: "Movie genres",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var companyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Company",
		Description: "Movie production company",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"logo_path": &graphql.Field{
				Type: graphql.String,
			},
			"origin_country": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var countryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Country",
		Description: "Movie production country",
		Fields: graphql.Fields{
			"iso_3166_1": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var pubUrlType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "PublicUrl",
		Description: "Movie screen shots public urls",
		Fields: graphql.Fields{
			"thumb": &graphql.Field{
				Type: graphql.String,
			},
			"full": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var screenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "ScreenShot",
		Description: "Movie screen shots",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"order": &graphql.Field{
				Type: graphql.Int,
			},
			"public_urls": &graphql.Field{
				Type: pubUrlType,
			},
		},
	},
)

var languageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Language",
		Description: "Movie spoken languages",
		Fields: graphql.Fields{
			"iso_639_1": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var movieType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Movie",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"adult": &graphql.Field{
				Type: graphql.Boolean,
			},
			"backdrop_Path": &graphql.Field{
				Type: graphql.String,
			},
			"belongs_To_Collection": &graphql.Field{
				Type: collectionType,
			},
			"budget": &graphql.Field{
				Type: graphql.Int,
			},
			"created_at": &graphql.Field{
				Type: graphql.Float,
			},
			"direct_link": &graphql.Field{
				Type: graphql.String,
			},
			"genres": &graphql.Field{
				Type: graphql.NewList(genreType),
			},
			"homepage": &graphql.Field{
				Type: graphql.String,
			},
			"imdb_id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"original_language": &graphql.Field{
				Type: graphql.String,
			},
			"original_title": &graphql.Field{
				Type: graphql.String,
			},
			"overview": &graphql.Field{
				Type: graphql.String,
			},
			"popularity": &graphql.Field{
				Type: graphql.Float,
			},
			"poster_path": &graphql.Field{
				Type: graphql.String,
			},
			"production_companies": &graphql.Field{
				Type: graphql.NewList(companyType),
			},
			"production_countries": &graphql.Field{
				Type: graphql.NewList(countryType),
			},
			"release_date": &graphql.Field{
				Type: graphql.String,
			},
			"revenue": &graphql.Field{
				Type: graphql.Int,
			},
			"runtime": &graphql.Field{
				Type: graphql.Int,
			},
			"screens": &graphql.Field{
				Type: graphql.NewList(screenType),
			},
			"spoken_languages": &graphql.Field{
				Type: graphql.NewList(languageType),
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"tagline": &graphql.Field{
				Type: graphql.String,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"tmdb_id": &graphql.Field{
				Type: graphql.Int,
			},
			"user_uid": &graphql.Field{
				Type: graphql.String,
			},
			"video": &graphql.Field{
				Type: graphql.Boolean,
			},
			"vote_average": &graphql.Field{
				Type: graphql.Float,
			},
			"vote_count": &graphql.Field{
				Type: graphql.Int,
			},
		},
		Description: "All about the movie",
	},
)

var tmdbType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "TMDB",
		Description: "Tmdb movie or tv show",
		Fields: graphql.Fields{
			"poster_path": &graphql.Field{
				Type: graphql.String,
			},
			"popularity": &graphql.Field{
				Type: graphql.Float,
			},
			"vote_count": &graphql.Field{
				Type: graphql.Int,
			},
			"video": &graphql.Field{
				Type: graphql.Boolean,
			},
			"media_type": &graphql.Field{
				Type: graphql.String,
			},
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"adult": &graphql.Field{
				Type: graphql.Boolean,
			},
			"backdrop_path": &graphql.Field{
				Type: graphql.String,
			},
			"original_language": &graphql.Field{
				Type: graphql.String,
			},
			"original_title": &graphql.Field{
				Type: graphql.String,
			},
			"genre_ids": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"vote_average": &graphql.Field{
				Type: graphql.Float,
			},
			"overview": &graphql.Field{
				Type: graphql.String,
			},
			"release_date": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
