package main

import (
	"github.com/joho/godotenv"
	"github.com/sergushka/ahoi-kino/log"
	"os"
	"strconv"
)

var (
	PORT      = 0
	CERT_FILE = ""
	KEY_FILE  = ""
	logger    = log.GetLogger()
)

func Load() {
	var err error

	if os.Getenv("PROD") == "" {
		err = godotenv.Load()
	}

	if err != nil {
		logger.Println(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		logger.Println(err)
		PORT = 9000
	}

	CERT_FILE = os.Getenv("CERT_FILE_PATH")
	KEY_FILE = os.Getenv("KEY_FILE_PATH")
}
