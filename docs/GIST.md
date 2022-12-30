# GIST

GitHub Gists sync

```golang
import (
	"context"
	"fmt"

	"github.com/guionardo/go-gstools/gist"
)

func main() {
	var (
		err error
		ctx context.Context
	)

	// Create a local gist instance from files in a folder
	myGist, err := gist.CreateGistFromLocalFolder("My gist", "/home/user/gist")
	if err != nil {
		panic(err)
	}

	// Create a new git context, using a github token
	ctx, err = gist.NewGitContext("github_token", context.Background())
	if err != nil {
		panic(err)
	}

	// Synchronize the local gist with the github gist
	remoteGist, action, err := gist.SyncGistFiles(ctx, myGist, "/home/user/gist")
	if err != nil {
		panic(err)
	}
	// Action can be Download, Upload or NoAction
	fmt.Printf("Action: %v\n", action)

	// Print the remote gist ID
	fmt.Printf("Remote gist ID: %s\n", *remoteGist.ID)
}
```