package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func LoadRoutes() []Route {
	routes := movieRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range LoadRoutes() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}
