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
	}, {
		Uri:     "/movies/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetMovie,
	}, {
		Uri:     "/movies",
		Method:  http.MethodPost,
		Handler: controllers.AddMovie,
	}, {
		Uri:     "/movies/{id}",
		Method:  http.MethodPut,
		Handler: controllers.EditMovie,
	}, {
		Uri:     "/movies/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteMovie,
	},
}
