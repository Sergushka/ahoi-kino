package graph

import (
	"github.com/graphql-go/graphql"
)

var movieFields = graphql.Fields{
	"movies": &graphql.Field{
		Type:        graphql.NewList(movieType),
		Description: "Get Movies specified by movieCount per page",
		Args: graphql.FieldConfigArgument{
			"movieCount": &graphql.ArgumentConfig{
				Type:         graphql.NewNonNull(graphql.Int),
				DefaultValue: 4,
				Description:  "How much movies per page",
			},
			"page": &graphql.ArgumentConfig{
				Type:         graphql.NewNonNull(graphql.Int),
				DefaultValue: 1,
				Description:  "Which page return movies for",
			},
		},
		Resolve: controllers.GetGraphQLMovies,
	},
	"movie": &graphql.Field{
		Type:        movieType,
		Description: "Get movie by Id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
				Description:  "Movie id in database",
			},
			"directLink": &graphql.ArgumentConfig{
				Type:         graphql.String,
				DefaultValue: "",
				Description:  "Movie directLink from tmdb",
			},
		},
		Resolve: controllers.GetGraphQLMovie,
	},
}

var movieMutationFields = graphql.Fields{
	"updateMovie": &graphql.Field{
		Type:        movieType,
		Description: "Update movie by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"adult": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"backdrop_Path": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"belongs_To_Collection": &graphql.ArgumentConfig{
				Type: collectionType,
			},
			"budget": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"created_at": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"direct_link": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"genres": &graphql.ArgumentConfig{
				Type: graphql.NewList(genreType),
			},
			"homepage": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"imdb_id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"original_language": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"original_title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"overview": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"popularity": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"poster_path": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"production_companies": &graphql.ArgumentConfig{
				Type: graphql.NewList(companyType),
			},
			"production_countries": &graphql.ArgumentConfig{
				Type: graphql.NewList(countryType),
			},
			"release_date": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"revenue": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"runtime": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"screens": &graphql.ArgumentConfig{
				Type: graphql.NewList(screenType),
			},
			"spoken_languages": &graphql.ArgumentConfig{
				Type: graphql.NewList(languageType),
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"tagline": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"title": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"tmdb_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"user_uid": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"video": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"vote_average": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"vote_count": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: controllers.UpdateGraphQLMovie,
	},
}
