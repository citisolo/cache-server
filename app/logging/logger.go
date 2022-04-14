package logging

import (
	log "github.com/sirupsen/logrus"
)

type AppLogs interface {
	Log(message string)
}

type AppLogger struct {
	Logger   *log.Logger
}

func NewAppLogger() AppLogs {
	return &AppLogger{
		Logger: log.New(),
	}
}

func (l *AppLogger) Log(message string) {
	fields := log.Fields{
		"App": "Info",
	}

	l.Logger.
		WithFields(fields).
		Info(message)
}