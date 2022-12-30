package gist

import (
	"os"
	"testing"
)

func getTokenOrSkipTests(t *testing.T) string {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		content, _ := os.ReadFile(".github_token")
		if len(content) > 0 {
			return string(content)
		}

		t.Skip("No token provided [env GITHUB_TOKEN, file .github_token], skipping tests")
		return ""
	}

	return token
}

