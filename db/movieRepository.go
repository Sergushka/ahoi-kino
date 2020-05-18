package db

import (
	"context"
	"github.com/sergushka/ahoi-kino/model"
)

type MovieRepository interface {
	AddMovie(movie *model.Movie) (*model.Movie, error)
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
		app.logger.Println("Fuck this shit")
	}

	return movie, nil
}
