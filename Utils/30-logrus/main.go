package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("beg")

	logFields := logrus.Fields{
		"details": logrus.Fields{
			"app_name":    "Some name",
			"app_version": "Some version",
		},
	}

	logger := logrus.New()
	logger.WithFields(logFields).Info("stevepro")

	logrus.Info("end")
}
