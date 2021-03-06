package db

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	lg "github.com/sergushka/ahoi-kino/log"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type FireApp struct {
	app    *firebase.App
	logger *log.Logger
}

const movieCollectionName = "movies"

var (
	fireApp *FireApp
	once    sync.Once
	logger  = lg.GetLogger()
)

func initializeAppWithServiceAccount() *firebase.App {
	dat, err := ioutil.ReadFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	if err != nil {
		logger.Fatalf("error reading file : %v\n", err)
	}

	opt := option.WithCredentialsJSON(dat)

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		logger.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func (fA *FireApp) GetDB() *firestore.Client {
	if fA == nil {
		logger.Fatal("FireApp is nil")
	}

	app, err := fA.app.Firestore(context.Background())

	if err != nil {
		logger.Fatalf("error initializing Firestore: %v\n", err)
	}

	return app
}

func GetApp() *FireApp {
	return fireApp
}

func NewApp() *FireApp {
	once.Do(func() {
		fireApp = &FireApp{
			logger: logger,
			app:    initializeAppWithServiceAccount(),
		}
	})
	return fireApp
}
