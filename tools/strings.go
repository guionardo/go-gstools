package tools

import "strings"

// JustNumbers returns only numbers from a string
func JustNumbers(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, str)
}
