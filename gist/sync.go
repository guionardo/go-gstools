package gist

import (
	"context"
	"errors"

	"github.com/google/go-github/v48/github"
)

func mapEqual(map1 map[github.GistFilename]github.GistFile, map2 map[github.GistFilename]github.GistFile) bool {
	if len(map1) != len(map2) {
		Log("map1 len: %d, map2 len: %d", len(map1), len(map2))
		return false
	}
	for key, value := range map1 {
		value2, ok := map2[key]
		if !ok {
			Log("missing key: %s in map2", key)
			return false
		}
		if value.Content == nil || value2.Content == nil {
			Log("missing content value.Content=%v, value2.Content=%v", value.Content, value2.Content)
			return false
		}
		if *value.Content != *value2.Content {
			Log("content not equal: %s != %s", *value.Content, *value2.Content)
			return false
		}
	}
	return true
}

func AreGistsFilesEquals(gist1 *github.Gist, gist2 *github.Gist) bool {
	if gist1 == nil || gist2 == nil || gist1.Files == nil || gist2.Files == nil {
		return false
	}
	return mapEqual(gist1.Files, gist2.Files) && mapEqual(gist2.Files, gist1.Files)
}

func SyncGistFiles(ctx context.Context, gist *github.Gist, localFolder string) (remoteGist *github.Gist, action SyncAction, err error) {
	action = NoAction
	if gist == nil {
		err = errors.New("Cannot sync empty gist")
		return
	}

	if gist.ID != nil {
		remoteGist, err = GetGistById(ctx, *gist.ID)
	} else if gist.Description != nil {
		remoteGist, err = GetGistByDescription(ctx, *gist.Description)
	} else {
		err = errors.New("Cannot sync gist without id or description")
		return
	}
	if remoteGist != nil && len(remoteGist.Files) > 0 {
		remoteGist, err = GetGistById(ctx, *remoteGist.ID)
	}

	action, err = getAction(localFolder, gist, remoteGist)
	if err != nil || action == NoAction {
		return
	}
	logs := getFilesDiff(gist, remoteGist)
	for _, log := range logs {
		Log(log)
	}
	switch action {
	case Download:
		err = SaveGistFilesToLocal(remoteGist, localFolder)

	case Upload:
		var client *github.Client
		client, err = GetClientFromContext(ctx)
		if err == nil {
			if remoteGist == nil {
				remoteGist, _, err = client.Gists.Create(ctx, gist)
			} else {
				remoteGist, _, err = client.Gists.Edit(ctx, *remoteGist.ID, gist)
			}
		}
	}

	return
}
