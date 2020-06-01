package tmdb

import (
	"encoding/json"
	"fmt"
	"github.com/sergushka/ahoi-kino/log"
	"github.com/sergushka/ahoi-kino/model"
	"net/http"
	"os"
	"strings"
)

const Space = "%20"

var (
	logger = log.GetLogger()
)

func GetTMDBMovies(movieName string) ([]model.TMDBMovieResponse, error) {
	movieName = strings.Join(strings.Fields(movieName), Space)
	apiKey := os.Getenv("API_KEY")
	host := os.Getenv("HOST")
	requestUrl := fmt.Sprintf("%s/3/search/multi?api_key=%s&language=en-US&query=%s&page=1&include_adult=false", host, apiKey, movieName)

	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	var movieResponse model.MovieResponse

	err = json.NewDecoder(resp.Body).Decode(&movieResponse)
	defer resp.Body.Close()

	if err != nil {
		logger.Printf("%T\n%s\n%#v\n", err, err, err)
		return nil, err
	}

	return movieResponse.Results, nil
}
