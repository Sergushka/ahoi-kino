package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/sergushka/ahoi-kino/model"
)

const statsRefId = "--stats--"

type MovieRepository interface {
	AddMovie(movie *model.Movie) (*model.Movie, error)
	GetMovies(moviesRequest model.MoviesRequest) ([]model.Movie, error)
	GetMovieById(id string) (model.Movie, error)
	GetMovieByDirectLink(directLink string) (model.Movie, error)
	UpdateMovieById(id string, update map[string]interface{}) (model.Movie, error)
}

type repo struct{}

func NewRepository() MovieRepository {
	return &repo{}
}

func (*repo) AddMovie(movie *model.Movie) (*model.Movie, error) {
	ctx := context.Background()
	app := GetApp()
	client := app.GetDB()

	collection := client.Collection(movieCollectionName)

	defer client.Close()

	_, _, err := collection.Add(ctx, map[string]interface{}{
		"Id":   movie.Id,
		"Text": movie.Title,
	})

	if err != nil {
		app.logger.Printf("Fuck this shit %v", err)
		return nil, err
	}

	return movie, nil
}

func (*repo) GetAllMovies() (movies []model.Movie, err error) {
	ctx := context.Background()
	app := GetApp()
	client := app.GetDB()

	collection := client.Collection(movieCollectionName)

	iter := collection.Documents(ctx)
	docs, err := iter.GetAll()

	if err != nil {
		return
	}

	defer client.Close()

	movies = make([]model.Movie, len(docs))
	var movie model.Movie

	i := 0
	for _, doc := range docs {
		if doc.Ref.ID == statsRefId {
			continue
		}
		err := doc.DataTo(&movie)
		if err != nil {
			logger.Println(err)
			continue
		}
		movie.Id = doc.Ref.ID
		movies[i] = movie
		i++
	}

	logger.Printf("found movies %v", len(movies))

	return
}

func (*repo) GetMovies(moviesRequest model.MoviesRequest) (movies []model.Movie, err error) {
	ctx := context.Background()
	app := GetApp()
	client := app.GetDB()

	collection := client.Collection(movieCollectionName)

	offset := (moviesRequest.Page - 1) * moviesRequest.MovieCount

	//offset + 1, first in collection is Stats not a movie
	iter := collection.Offset(offset + 1).Limit(moviesRequest.MovieCount).Documents(ctx)
	docs, err := iter.GetAll()

	if err != nil {
		return
	}

	defer client.Close()

	movies = make([]model.Movie, len(docs))
	var movie model.Movie

	i := 0
	for _, doc := range docs {
		if doc.Ref.ID == statsRefId {
			continue
		}
		err := doc.DataTo(&movie)
		if err != nil {
			logger.Println(err)
			continue
		}
		movie.Id = doc.Ref.ID
		movies[i] = movie
		i++
	}

	logger.Printf("found movies %v", len(movies))

	return
}

func (*repo) GetMovieById(id string) (movie model.Movie, err error) {
	ctx := context.Background()
	app := GetApp()
	client := app.GetDB()

	collection := client.Collection(movieCollectionName)

	doc, err := collection.Doc(id).Get(ctx)

	if err != nil {
		logger.Printf("movie with id [%s] not found", id)
		return
	}

	defer client.Close()

	err = doc.DataTo(&movie)

	if err != nil {
		logger.Printf("Can't convert %v because %v", doc.Ref.Path, err)
		return
	}

	movie.Id = doc.Ref.ID
	logger.Printf("found movie %s By Id %s", movie.Name, movie.Id)

	return
}

func (*repo) GetMovieByDirectLink(directLink string) (movie model.Movie, err error) {
	ctx := context.Background()
	app := GetApp()
	client := app.GetDB()

	collection := client.Collection(movieCollectionName)

	docs, err := collection.Where("directLink", "==", directLink).Documents(ctx).GetAll()

	if err != nil {
		logger.Printf("movie with directLink [%s] %v", directLink, err)
		return
	}

	defer client.Close()

	//Only one with unique direct link, if none found it will be error
	doc := docs[0]

	err = doc.DataTo(&movie)

	if err != nil {
		logger.Printf("Can't convert %v because %v", doc.Ref.Path, err)
		return
	}

	movie.Id = doc.Ref.ID
	logger.Printf("found movie %s by directLink", movie.Name)

	return
}

func (*repo) UpdateMovieById(id string, update map[string]interface{}) (movie model.Movie, err error) {
	ctx := context.Background()
	app := GetApp()
	client := app.GetDB()

	collection := client.Collection(movieCollectionName)

	document := collection.Doc(id)

	defer client.Close()

	_, err = document.Set(ctx, update, firestore.MergeAll)

	if err != nil {
		logger.Printf("Can't update movie %v", err)
		return
	}

	doc, err := document.Get(ctx)
	if err != nil {
		logger.Printf("Can't get movie %v", err)
		return
	}

	err = doc.DataTo(&movie)

	if err != nil {
		logger.Printf("Can't convert movie %v", err)
		return
	}

	movie.Id = id
	logger.Printf("updated movie %s", movie.Name)

	return
}
