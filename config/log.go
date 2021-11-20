package config

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func SetupLog() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(io.MultiWriter(os.Stdout))
	log.SetOutput(logger.Writer())

	return logger
}