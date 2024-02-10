package common

import (
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func NewLogger(stage string) *logrus.Entry {
	var log = logrus.New()

	if debug := os.Getenv("debug"); debug == "true" {
		log.Level = logrus.DebugLevel
	} else {
		log.Level = logrus.InfoLevel
	}

	return log.WithFields(logrus.Fields{"run_id": uuid.New(), "stage": stage})
}
