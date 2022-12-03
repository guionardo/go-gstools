package logger

import (
	"bytes"
	"log"
	"runtime"
	"testing"
)

func TestLogger_Logging(t *testing.T) {

	t.Run("Default", func(t *testing.T) {
		logBuffer := &bytes.Buffer{}
		internalLogger := log.New(logBuffer, "TEST", 0)
		logger := &Logger{}
		logger.SetLogger(internalLogger).SetDebug(true)
		logger.Debugf("Test %s", "Debug")
		logger.Infof("Test %s", "Info")
		logger.Warningf("Test %s", "Warning")
		logger.Errorf("Test %s", "Error")
		var expected string
		if runtime.GOOS == "windows" {
			expected = "TEST[DEBUG] Test Debug\nTEST[INFO] Test Info\nTEST[WARNING] Test Warning\nTEST[ERROR] Test Error\n"
		} else {
			expected = "TEST\x1b[36m[DEBUG] Test Debug\x1b[0m\nTEST\x1b[32m[INFO] Test Info\x1b[0m\nTEST\x1b[33m[WARNING] Test Warning\x1b[0m\nTEST\x1b[31m[ERROR] Test Error\x1b[0m\n"
		}
		got := logBuffer.String()
		if got != expected {
			t.Errorf("Logger.Debug() = %v, want %v", got, "TEST [DEBUG] Test Debug")
		}

	})

}

func TestLogger_AutoSetup(t *testing.T) {
	type myStruct struct {
		Logger
		data int
	}

	t.Run("Default", func(t *testing.T) {
		m := &myStruct{}
		m.AutoSetup("myStruct")
		m.Infof("Test %s", "Info")
		if m.debug {
			t.Errorf("Logger.debug = %v, want %v", m.debug, false)
		}
		if m.logger.Prefix() != "myStruct " {
			t.Errorf("Logger.logger.Prefix() = %v, want %v", m.logger.Prefix(), "myStruct ")
		}
	})

}

func TestLogger_Default(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		Default().Infof("Test %s", "Info")
		if Default().debug {
			t.Errorf("Logger.debug = %v, want %v", Default().debug, false)
		}
	})
}
