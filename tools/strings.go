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
func JoinNotEmpty(s ...string) (r string) {
	r = ""
	for _, v := range s {
		if len(v) == 0 {
			continue
		}
		if len(r) > 0 {
			r += ", "
		}
		r += v
	}
	return
}