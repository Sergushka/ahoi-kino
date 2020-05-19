package controllers

import (
	"encoding/json"
	"github.com/sergushka/ahoi-kino/db"
	"github.com/sergushka/ahoi-kino/model"
	"github.com/sergushka/ahoi-kino/utils"
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	database := db.NewTestRepository()
	movies, err := database.GetMovies()

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Result: movies,
	}

	js, err := json.Marshal(response)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Single movie"))
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("New movie added"))
}

func EditMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Movie edited"))
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Movie deleted"))
}
