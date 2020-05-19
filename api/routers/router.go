package routers

import (
	"github.com/gorilla/mux"
	"github.com/sergushka/ahoi-kino/api/routers/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutes(r)
}
