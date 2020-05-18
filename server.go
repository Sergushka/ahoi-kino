package main

import (
	"github.com/gorilla/mux"
	"github.com/sergushka/ahoi-kino/db"
	main2 "github.com/sergushka/ahoi-kino/log"
	"github.com/sergushka/ahoi-kino/services"
	"net/http"
)

func main() {
	logger := main2.GetLogger()
	router := mux.NewRouter()

	db.NewApp()

	const port = ":8080"

	//database := db.NewTestRepository()

	//_, _ = database.AddMovie(test)

	movie := services.GetMovieByName("hitman")

	logger.Println(movie)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("got %v", r.Header)
	})

	logger.Println("server starting...")
	err := http.ListenAndServe(port, router)

	if err != nil {
		logger.Fatalf("server failed to start... %v", err)
	}
}
