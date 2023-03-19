package utils

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger
var once sync.Once

func init() {
	once.Do(func() {
		Logger = log.New()
		Logger.SetFormatter(&log.JSONFormatter{})
		Logger.SetLevel(log.InfoLevel)
	})
}
