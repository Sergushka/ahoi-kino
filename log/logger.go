package log

import (
	"log"
	"os"
)

func GetLogger() *log.Logger {
	return log.New(os.Stdout, "ahoi ", log.LstdFlags|log.Lshortfile)
}
