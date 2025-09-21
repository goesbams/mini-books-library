package utils

import "github.com/sirupsen/logrus"

func InitializeLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return logger
}
