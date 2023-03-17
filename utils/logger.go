package utils

import (
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New()
	Logger.SetFormatter(&log.JSONFormatter{})
	Logger.SetLevel(log.InfoLevel)
}
