package logger

import log "github.com/sirupsen/logrus"

func GetLogger(layer, functionName string) *log.Logger {
	logger := log.WithFields(log.Fields{
		"layer":        layer,
		"functionName": functionName,
	}).Logger

	logger.SetFormatter(&log.JSONFormatter{})

	return logger
}
