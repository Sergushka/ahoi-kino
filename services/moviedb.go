package services

import (
	"encoding/json"
	"github.com/sergushka/ahoi-kino/log"
	"github.com/sergushka/ahoi-kino/model"
	"net/http"
	"strings"
)

var (
	logger = log.GetLogger()
)

func GetMovieByName(name string) *model.Movie {
	name = strings.Join(strings.Fields(name), "%20")
	requestUrl := "https://api.themoviedb.org/3/search/multi?api_key=" + ApiKey + "&language=en-US&query=" + name + "&page=1&include_adult=false"

	resp, err := http.Get(requestUrl)

	if err != nil {
		logger.Fatalf("No movie found %v", err)
	}

	defer resp.Body.Close()

	var m = new(model.MovieResponse)

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&m)

	if err != nil {
		logger.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	return &m.Results[0]
}
