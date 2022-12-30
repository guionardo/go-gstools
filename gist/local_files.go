package gist

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/google/go-github/v48/github"
)

func readFile(fileName string) (gist github.GistFile, err error) {
	defer Log("Reading file %s - %v", fileName, err)
	content, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	mimeType, language := detectContentTypeAndLanguage(fileName, content)
	gist = github.GistFile{
		Content:  github.String(string(content)),
		Size:     github.Int(len(content)),
		Filename: github.String(path.Base(fileName)),
		Language: github.String(language),
		Type:     github.String(mimeType),
	}
	return
}

func CreateGistFromLocalFolder(description string, localFolder string) (gist *github.Gist, err error) {
	files, err := ioutil.ReadDir(localFolder)
	if err != nil {
		return
	}
	fileNames := make([]string, 0, len(files))
	shortFileNames := make([]string, 0, len(files))
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, path.Join(localFolder, file.Name()))
			shortFileNames = append(shortFileNames, file.Name())
		}
	}
	if len(fileNames) == 0 {
		return nil, fmt.Errorf("No files in folder %s", localFolder)
	}
	Log("Getting files from folder %s: %v", localFolder, shortFileNames)
	gist, err = CreateGistFromLocalFiles(description, fileNames...)

	return

}
func CreateGistFromLocalFiles(description string, files ...string) (gist *github.Gist, err error) {
	gist = &github.Gist{
		Description: github.String(description),
		Public:      github.Bool(true),
		Files:       make(map[github.GistFilename]github.GistFile),
	}
	lastUpdate := time.Time{}
	for _, file := range files {
		stat, err := os.Stat(file)
		if err != nil || stat.IsDir() {
			continue
		}
		if lastUpdate.Before(stat.ModTime()) {
			lastUpdate = stat.ModTime()
		}
		gistFile, err := readFile(file)

		if err == nil {
			gist.Files[github.GistFilename(path.Base(file))] = gistFile
		}
	}
	gist.UpdatedAt = &lastUpdate
	return gist, nil
}

func SaveGistFilesToLocal(gist *github.Gist, localFolder string) (err error) {
	for fileName, file := range gist.Files {
		localFileName := path.Join(localFolder, string(fileName))
		if err = os.WriteFile(localFileName, []byte(*file.Content), 0644); err == nil {
			err = os.Chtimes(localFileName, *gist.CreatedAt, *gist.UpdatedAt)
		}
		if err != nil {
			break
		}
	}
	return
}
