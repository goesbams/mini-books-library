package utils

import "github.com/sirupsen/logrus"

func initializeLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logger
}
