package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sergushka/ahoi-kino/log"
	"github.com/sergushka/ahoi-kino/model"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const Space = "%20"

var (
	logger = log.GetLogger()
)

func GetMovieByName(movieName string) (*model.Movie, error) {
	movieName = strings.Join(strings.Fields(movieName), Space)
	apiKey := os.Getenv("API_KEY")
	host := os.Getenv("HOST")
	requestUrl := fmt.Sprintf("%s/3/search/multi?api_key=%s&language=en-US&query=%s&page=1&include_adult=false", host, apiKey, movieName)

	logger.Println(requestUrl)

	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	rb, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var movieResponse model.MovieResponse

	err = json.Unmarshal(rb, &movieResponse)

	if err != nil {
		logger.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	if len(movieResponse.Results) == 0 {
		return nil, errors.New(fmt.Sprintf("movie %s not found", movieName))
	}

	return &movieResponse.Results[0], nil
}
