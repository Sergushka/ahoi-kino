package main

import (
	"fmt"
	"github.com/sergushka/ahoi-kino/api/routers"
	"github.com/sergushka/ahoi-kino/db"
	"net/http"
	"time"
)

func Run() {
	Load()
	db.NewApp()

	listen()
}

func listen() {
	router := routers.New()

	logger.Printf("Listening on %v", PORT)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      router,
	}

	logger.Fatal(srv.ListenAndServe())
}
