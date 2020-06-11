package routes

import (
	"github.com/sergushka/ahoi-kino/api/controllers"
	"net/http"
)

var movieRoutes = []Route{
	{
		Uri:     "/movies",
		Method:  http.MethodGet,
		Handler: controllers.GetMovies,
	},
	{
		Uri:     "/movies/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetMovie,
	},
	{
		Uri:     "/graphql",
		Method:  http.MethodGet,
		Handler: controllers.GraphQL,
	},
	{
		Uri:     "/graphql",
		Method:  http.MethodPost,
		Handler: controllers.GraphQLPost,
	},
}
