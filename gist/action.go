package gist

import (
	"errors"
	"fmt"

	"github.com/google/go-github/v48/github"
)

type SyncAction byte

const (
	Upload SyncAction = iota
	Download
	NoAction
)

func getAction(localFolder string, localGist *github.Gist, remoteGist *github.Gist) (action SyncAction, err error) {
	action = NoAction
	logMsg := "No action"
	if localGist == nil && remoteGist == nil {
		err = errors.New("Cannot sync empty gist")
	} else if remoteGist == nil {
		action = Upload
		logMsg = "Upload due to no remote gist"
	} else if !AreGistsFilesEquals(localGist, remoteGist) {
		if localGist == nil {
			logMsg = "Download due to no local gist"
			action = Download
		} else if remoteGist == nil {
			logMsg = "Upload due to no remote gist"
			action = Upload
		} else if localGist.UpdatedAt.Before(*remoteGist.UpdatedAt) {
			logMsg = fmt.Sprintf("Download due to remote gist [%v] newer than local [%v]", remoteGist.UpdatedAt, localGist.UpdatedAt)
			action = Download
		} else if localGist.UpdatedAt.After(*remoteGist.UpdatedAt) {
			logMsg = fmt.Sprintf("Upload due to local gist [%v] newer than remote [%v]", localGist.UpdatedAt, remoteGist.UpdatedAt)
			action = Upload
		}
	}
	Log("Sync action: %s - %v", logMsg, err)
	return
}
