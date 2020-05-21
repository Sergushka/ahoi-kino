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
				Type:         graphql.Int,
				DefaultValue: 4,
				Description:  "How much movies per page",
			},
			"page": &graphql.ArgumentConfig{
				Type:         graphql.Int,
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
