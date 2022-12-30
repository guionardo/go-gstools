package gist

import (
	"context"
	"testing"
)

func TestNewGitContext(t *testing.T) {

	t.Run("Default", func(t *testing.T) {
		token := getTokenOrSkipTests(t)
		ctx := context.Background()
		got, err := NewGitContext(token, ctx)
		if err != nil {
			t.Errorf("NewGitContext() error = %v", err)
		}
		if _, err = GetClientFromContext(got); err != nil {
			t.Errorf("NewGitContext() %v", err)
		}
	})
}
