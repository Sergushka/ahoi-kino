package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/sergushka/ahoi-kino/api/graph"
	"github.com/sergushka/ahoi-kino/db"
	"github.com/sergushka/ahoi-kino/model"
	"github.com/sergushka/ahoi-kino/utils"
	"net/http"
	"path"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	database := db.NewRepository()

	var moviesRequest model.MoviesRequest
	err := json.NewDecoder(r.Body).Decode(&moviesRequest)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	movies, err := database.GetMovies(moviesRequest)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response := model.Response{
		Result: movies,
	}

	utils.JSON(w, http.StatusOK, response)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	database := db.NewRepository()
	id := path.Base(r.RequestURI)
	w.Header().Set("Content-Type", "application/json")

	if id == "" {
		utils.ERROR(w, http.StatusNoContent, errors.New("please specify id if /movies/{id}"))
		return
	}

	movie, err := database.GetMovieById(id)

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, errors.New(fmt.Sprintf("movie with id %s not found", id)))
		return
	}

	response := model.Response{
		Result: movie,
	}

	utils.JSON(w, http.StatusFound, response)
}

func GraphQL(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	params := graphql.Params{
		Schema:        graph.New(),
		RequestString: query,
	}
	result := graphql.Do(params)
	json.NewEncoder(w).Encode(result)
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
