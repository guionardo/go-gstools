package gist

import (
	"log"
	"os"
)

var defaultLogger *log.Logger

func SetLogger(logger *log.Logger) {
	defaultLogger = logger
}

func SetDefaultLogger() {
	defaultLogger = log.New(os.Stdout, "gist", log.LstdFlags)
}

func Log(format string, v ...interface{}) {
	if defaultLogger != nil {
		defaultLogger.Printf(format, v...)
	}
}
