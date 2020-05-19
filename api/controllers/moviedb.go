package controllers

import (
	"encoding/json"
	"errors"
	"github.com/sergushka/ahoi-kino/log"
	"github.com/sergushka/ahoi-kino/model"
	"net/http"
	"os"
	"strings"
)

var (
	logger = log.GetLogger()
)

func GetMovieByName(name string) (*model.Movie, error) {
	name = strings.Join(strings.Fields(name), "%20")
	apiKey := os.Getenv("API_KEY")

	requestUrl :=
		"https://api.themoviedb.org/3/search/multi?api_key=" +
			apiKey + "&language=en-US&query=" +
			name + "&page=1&include_adult=false"

	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var m = new(model.MovieResponse)

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&m)

	if err != nil {
		logger.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	if len(m.Results) == 0 {
		return nil, errors.New("movie " + name + " not found")
	}

	return &m.Results[0], nil
}
