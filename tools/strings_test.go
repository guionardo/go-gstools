package tools

import (
	"testing"
)

func TestJoinNotEmpty(t *testing.T) {
	tests := []struct {
		name  string
		s     []string
		wantR string
	}{
		{
			name:  "single arg",
			s:     []string{"ABCD"},
			wantR: "ABCD",
		},
		{
			name:  "one empty arg",
			s:     []string{"ABCD", "", "1234"},
			wantR: "ABCD, 1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := JoinNotEmpty(tt.s...); gotR != tt.wantR {
				t.Errorf("JoinNotEmpty() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestJustNumbers(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Numbers 1",
			arg:  " 12345 ",
			want: "12345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JustNumbers(tt.arg); got != tt.want {
				t.Errorf("JustNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
