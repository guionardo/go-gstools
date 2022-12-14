package gist

import (
	"context"
	"errors"

	"github.com/google/go-github/v48/github"
)

func NewGitContext(token string, ctx context.Context) (context.Context, error) {
	httpClient := getHttpClient(token, ctx)
	gitClient := github.NewClient(httpClient)
	_, response, err := gitClient.APIMeta(ctx)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 403 {
		return nil, UnauthorizedError
	} else if response.StatusCode >= 400 {
		return nil, HttpResponseError(response.StatusCode, response.Status)
	}

	return context.WithValue(ctx, "client", gitClient), nil
}

func GetClientFromContext(ctx context.Context) (*github.Client, error) {
	client := ctx.Value("client").(*github.Client)
	if client == nil {
		return nil, errors.New("No client in context")
	}
	return client, nil
}
