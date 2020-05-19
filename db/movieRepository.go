package db

import (
	"context"
	"github.com/sergushka/ahoi-kino/model"
)

const statsRefId = "--stats--"

type MovieRepository interface {
	AddMovie(movie *model.Movie) (*model.Movie, error)
	GetMovies() ([]model.Movie, error)
}

type repo struct{}

func NewTestRepository() MovieRepository {
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

func (*repo) GetMovies() (movies []model.Movie, err error) {
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

	return
}
