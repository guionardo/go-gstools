package logger

import "runtime"

const (
	Debug int = iota
	Info
	Warn
	Error
)

var (
	DebugColor = ""
	InfoColor  = ""
	WarnColor  = ""
	ErrorColor = ""
	ResetColor = ""
)

/*
*
Common Colors

	Reset = "\033[0m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Blue = "\033[34m"
	Purple = "\033[35m"
	Cyan = "\033[36m"
	Gray = "\033[37m"
	White = "\033[97m"
*/
func init() {
	if runtime.GOOS != "windows" {
		DebugColor = "\033[36m" // Cyan
		InfoColor = "\033[32m"  // Green
		WarnColor = "\033[33m"  // Yellow
		ErrorColor = "\033[31m" // Red
		ResetColor = "\033[0m"  // Reset
	}
}

func Colorize(text string, color int) string {
	switch color {
	case Error:
		return ErrorColor + text + ResetColor
	case Warn:
		return WarnColor + text + ResetColor
	case Info:
		return InfoColor + text + ResetColor
	case Debug:
		return DebugColor + text + ResetColor
	default:
		return text
	}
}
