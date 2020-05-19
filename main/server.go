package main

import (
	"fmt"
	"github.com/sergushka/ahoi-kino/api/routers"
	"github.com/sergushka/ahoi-kino/db"
	"net/http"
)

func Run() {
	db.NewApp()
	Load()

	//movie, _ := controllers.GetMovieByName("hitman")
	//
	//logger.Println(movie)

	listen()
}

func listen() {
	router := routers.New()
	logger.Printf("Listening on %v", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), router)

	if err != nil {
		logger.Fatalf("server failed to start... %v", err)
	}
}
