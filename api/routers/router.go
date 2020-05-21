package routers

import (
	"github.com/gorilla/mux"
	"github.com/sergushka/ahoi-kino/api/routers/routes"
	"net/http"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(commonMiddleware)
	return routes.SetupRoutes(r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
