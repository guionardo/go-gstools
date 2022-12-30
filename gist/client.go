package gist

import (
	"context"

	"net/http"

	"golang.org/x/oauth2"
)

var clientPool = make(map[string]*http.Client)

func getHttpClient(token string, ctx context.Context) *http.Client {
	if client, ok := clientPool[token]; ok {
		return client
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client := oauth2.NewClient(ctx, ts)
	clientPool[token] = client
	return client
}
