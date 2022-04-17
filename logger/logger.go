package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	Logger = log.New()
	Logger.SetLevel(log.DebugLevel)
	Logger.SetFormatter(&log.TextFormatter{})
	Logger.SetOutput(file)
}