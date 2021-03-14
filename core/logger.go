package core

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    &logrus.TextFormatter{},
		Level:        logrus.DebugLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
}

func GetLogger() *logrus.Logger {
	return logger
}
