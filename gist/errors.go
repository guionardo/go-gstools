package gist

import (
	"errors"
	"fmt"
)

var (
	UnauthorizedError         = errors.New("Unauthorized")
	EmptyGistError            = errors.New("Cannot post empty gist")
	EmptyDescriptionGistError = errors.New("Cannot get gist with empty description")
	EmptyIdGistError          = errors.New("Cannot get gist with empty id")
)

func HttpResponseError(statusCode int, status string) error {
	return fmt.Errorf("Http error %d - %s", statusCode, status)
}
