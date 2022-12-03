// Package logger implements type for adding logging capabilities to
// your application.
//
// The logger is a wrapper around the standard log package.
// It adds the ability to enable/disable debug logging and to colorize
// the output.

package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	logger   *log.Logger
	debug    bool
	colorize func(string, int) string
}

const (
	DEBUG   = "[DEBUG]"
	INFO    = "[INFO]"
	ERROR   = "[ERROR]"
	WARNING = "[WARNING]"
)

var defaultLogger *Logger

func Default() *Logger {
	if defaultLogger == nil {
		defaultLogger = &Logger{}
		defaultLogger.AutoSetup("Default")
	}
	return defaultLogger
}

func (r *Logger) SetLogger(logger *log.Logger) *Logger {
	r.logger = logger
	if r.colorize == nil {
		r.colorize = Colorize
	}
	return r
}

// Setup logger with default settings:
// - Debug logging disabled
// - Colorized output enabled
// - Output to stdout
// - Prefix "Type"
// - Ldate | Ltime | Lmsgprefix
func (r *Logger) AutoSetup(loggerName string) *Logger {
	loggerName = strings.TrimSpace(loggerName)
	if len(loggerName) > 0 {
		loggerName = loggerName + " "
	}
	internalLogging := log.New(os.Stdout, loggerName, log.Lmsgprefix|log.LstdFlags)
	return r.SetLogger(internalLogging).SetDebug(false).UseColors(true)
}

// Enable debug logging
func (r *Logger) SetDebug(debug bool) *Logger {
	r.debug = debug
	return r
}

// Activate colorization of the output. This is only supported on linux/macOS.
func (r *Logger) UseColors(enable bool) *Logger {
	if enable && runtime.GOOS != "windows" {
		r.colorize = Colorize
	} else {
		r.colorize = func(format string, _ int) string {
			return format
		}
	}
	return r
}

func (r *Logger) log(prefix string, color int, format string, a ...interface{}) {
	if r.logger == nil {
		return
	}
	if len(prefix) > 0 {
		format = prefix + " " + format
	}
	r.logger.Printf(r.colorize(fmt.Sprintf(format, a...), color))
}

func (r *Logger) Debugf(format string, a ...interface{}) {
	if r.debug {
		r.log(DEBUG, Debug, format, a...)
	}
}

func (r *Logger) Infof(format string, a ...interface{}) {
	r.log(INFO, Info, format, a...)
}

func (r *Logger) Errorf(format string, a ...interface{}) {
	r.log(ERROR, Error, format, a...)
}

func (r *Logger) Warningf(format string, a ...interface{}) {
	r.log(WARNING, Warn, format, a...)
}

func (r *Logger) Warnf(format string, a ...interface{}) {
	r.Warningf(format, a...)
}

func (r *Logger) Fatalf(format string, a ...interface{}) {
	r.logger.Fatalf(format, a...)
}
