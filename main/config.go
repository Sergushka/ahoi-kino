package main

import (
	"github.com/joho/godotenv"
	"github.com/sergushka/ahoi-kino/log"
	"os"
	"strconv"
)

var (
	PORT   = 0
	logger = log.GetLogger()
)

func Load() {
	var err error
	err = godotenv.Load()

	if err != nil {
		logger.Println(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		logger.Println(err)
		PORT = 9000
	}
}
