package gist

import (
	"context"
	"errors"

	"github.com/google/go-github/v48/github"
)

func GetGistById(ctx context.Context, gistId string) (*github.Gist, error) {
	if gistId == "" {
		return nil, EmptyIdGistError
	}
	client, err := GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	gist, _, err := client.Gists.Get(ctx, gistId)
	if err != nil {
		return nil, err
	}
	return gist, err
}

func GetGistByDescription(ctx context.Context, gistDescription string) (*github.Gist, error) {
	if gistDescription == "" {
		return nil, EmptyDescriptionGistError
	}
	client, err := GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}
	gists, _, err := client.Gists.ListAll(ctx, &github.GistListOptions{})
	if err != nil {
		return nil, err
	}
	for _, gist := range gists {
		if gist != nil && gist.Description != nil && *gist.Description == gistDescription {
			return gist, err
		}
	}
	return nil, err
}

func PostGist(ctx context.Context, gist *github.Gist) (*github.Gist, error) {
	if gist == nil {
		return nil, EmptyGistError
	}
	client, err := GetClientFromContext(ctx)
	if err != nil {
		return nil, err
	}

	gist, _, err = client.Gists.Create(ctx, gist)
	return gist, err
}

func DeleteGist(ctx context.Context, gistId string) error {
	if gistId == "" {
		return errors.New("Cannot delete empty gist")
	}
	client, err := GetClientFromContext(ctx)
	if err != nil {
		return err
	}
	_, err = client.Gists.Delete(ctx, gistId)
	return err
}
