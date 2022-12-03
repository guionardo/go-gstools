package logger

import (
	"testing"
)

func TestConsoleColor_Paint(t *testing.T) {
	const text = "Test"

	tests := []struct {
		name  string
		level int
		want  string
	}{

		{
			name:  "Error",
			level: Error,
			want:  "\033[31mTest\033[0m",
		},
		{
			name:  "Warn",
			level: Warn,
			want:  "\033[33mTest\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Colorize(text, tt.level); got != tt.want {
				t.Errorf("ConsoleColor.Paint() = %v, want %v", got, tt.want)
			}
		})
	}
}
