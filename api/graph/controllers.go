package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/sergushka/ahoi-kino/api/tmdb"
	"github.com/sergushka/ahoi-kino/db"
	"github.com/sergushka/ahoi-kino/log"
	"github.com/sergushka/ahoi-kino/model"
)

var (
	logger = log.GetLogger()
)

type GraphQLControllers interface {
	GetTMDBAll(p graphql.ResolveParams) (interface{}, error)
	GetGraphQLMovies(p graphql.ResolveParams) (interface{}, error)
	GetGraphQLMovie(p graphql.ResolveParams) (interface{}, error)
	UpdateGraphQLMovie(p graphql.ResolveParams) (interface{}, error)
}

var controllers = NewControllers()

type controller struct{}

func NewControllers() GraphQLControllers {
	return &controller{}
}

func (*controller) GetTMDBAll(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["name"].(string)

	return tmdb.GetTMDBMovies(name)
}

func (*controller) GetGraphQLMovies(p graphql.ResolveParams) (interface{}, error) {
	database := db.NewRepository()

	movieCount := p.Args["movieCount"]
	page := p.Args["page"]
	var moviesRequest model.MoviesRequest

	moviesRequest.MovieCount = movieCount.(int)
	moviesRequest.Page = page.(int)

	movies, err := database.GetMovies(moviesRequest)

	if err != nil {
		logger.Printf("Fuck %v", err)
		return nil, err
	}

	return movies, nil
}

func (*controller) GetGraphQLMovie(p graphql.ResolveParams) (interface{}, error) {
	database := db.NewRepository()

	directLink := p.Args["directLink"]
	id := p.Args["id"]

	var movie model.Movie
	var err error

	if id != "" {
		movie, err = database.GetMovieById(id.(string))
		if err != nil {
			logger.Printf("Fuck %v", err)
			return nil, err
		}
	} else {
		movie, err = database.GetMovieByDirectLink(directLink.(string))
		if err != nil {
			logger.Printf("Fuck %v", err)
			return nil, err
		}
	}

	return movie, nil
}

func (*controller) UpdateGraphQLMovie(p graphql.ResolveParams) (interface{}, error) {
	database := db.NewRepository()

	id := p.Args["id"].(string)

	movie, err := database.UpdateMovieById(id, p.Args)

	if err != nil {
		logger.Printf("Fuck %v", err)
		return nil, err
	}

	return movie, nil
}
