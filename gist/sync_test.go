package gist

import (
	"context"
	"os"
	"path"
	"testing"
	"time"
)

func createTestFiles(folder string, fileTime time.Time, changed bool) []string {
	createFile := func(filename string, content string) string {
		file := path.Join(folder, filename)
		os.WriteFile(file, []byte(content), 0644)
		os.Chtimes(file, fileTime, fileTime)
		return file
	}
	endFile := ""
	if changed {
		endFile = "\n\n"
	}
	files := []string{
		createFile("file1.txt", "file1"+endFile),
		createFile("file2.md", "# TITLE\n\nfile2"+endFile),
		createFile("file3.go", "package main\n\nfunc main() {\n\tprintln(\"file3\")\n}"+endFile),
		createFile("file4.py", "print(\"file4\")"+endFile),
		createFile("file5.js", "console.log(\"file5\")"+endFile),
	}

	return files
}

func TestCreateGistFromLocalFiles(t *testing.T) {
	SetDefaultLogger()
	t.Run("Default", func(t *testing.T) {
		files := createTestFiles(t.TempDir(), time.Now(), false)

		gotGist, err := CreateGistFromLocalFiles("Test gist", files...)
		if err != nil {
			t.Errorf("CreateGistFromLocalFiles() error = %v", err)
			return
		}
		if len(gotGist.Files) != len(files) {
			t.Errorf("CreateGistFromLocalFiles() = %v, want %v", len(gotGist.Files), len(files))
		}

	})
}

func TestSyncGistFiles(t *testing.T) {
	SetDefaultLogger()
	const testGist = "Test GIST"

	tmpDir1 := t.TempDir()
	files1 := createTestFiles(tmpDir1, time.Now().Add(-time.Hour), false)

	tmpDir2 := t.TempDir()
	files2 := createTestFiles(tmpDir2, time.Now(), true)

	tmpDir3 := t.TempDir()
	files3 := createTestFiles(tmpDir3, time.Now().Add(-time.Minute), false)

	cases := []struct {
		name                string
		folder              string
		files               []string
		remoteGistMustExist bool
		deleteOnExit        bool
		action              SyncAction
		startup             bool
	}{
		{"Startup [clear]", "", nil, false, false, NoAction, true},
		{"Default [Upload]", tmpDir1, files1, false, false, Upload, false},
		{"No modified [NoAction]", tmpDir1, files1, true, false, NoAction, false},
		{"Modified [Download]", tmpDir2, files2, true, false, Download, false},
		{"Modified [Upload]", tmpDir3, files3, true, true, NoAction, false},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			token := getTokenOrSkipTests(t)
			ctx, err := NewGitContext(token, context.Background())
			if err != nil {
				t.Errorf("NewGitContext() error = %v", err)
				return
			}
			if tt.startup {
				t.Log("Clearing all gists")
				unneededGist, _ := GetGistByDescription(ctx, testGist)
				if unneededGist != nil {
					DeleteGist(ctx, *unneededGist.ID)
				}
				return
			}
			localGist, err := CreateGistFromLocalFolder(testGist, tt.folder)
			if err != nil {
				t.Errorf("CreateGistFromLocalFiles(%s) error = %v", tt.folder, err)
				return
			}
			if len(localGist.Files) != len(tt.files) {
				t.Errorf("CreateGistFromLocalFiles(%s) = %v, want %v", tt.folder, len(localGist.Files), len(tt.files))
				return
			}
			remoteGist, action, err := SyncGistFiles(ctx, localGist, tt.folder)
			if err != nil {
				t.Errorf("SyncGistFiles() error = %v", err)
				return
			}
			if tt.remoteGistMustExist && remoteGist == nil {
				t.Errorf("Expected remote gist exists")
				return
			}
			if tt.action != action {
				t.Errorf("SyncGistFiles() action = %v, want %v", action, tt.action)
				return
			}
			updatedGist, _ := CreateGistFromLocalFolder(testGist, tt.folder)
			if !AreGistsFilesEquals(updatedGist, remoteGist) {
				t.Errorf("SyncGistFiles() = %v, want %v", localGist, remoteGist)
			}
			if remoteGist != nil && tt.deleteOnExit {
				DeleteGist(ctx, *remoteGist.ID)
			}
		})
	}

}
