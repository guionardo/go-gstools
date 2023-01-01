package gist

import (
	"os"
	"reflect"
	"testing"
)

func Test_getFilesDiff(t *testing.T) {
	tmpDir := t.TempDir()

	file := tmpDir + "/file"
	os.WriteFile(file, []byte("file1\nline 2\nend of file"), 0644)
	gist1, _ := CreateGistFromLocalFiles("Test 1", file)

	os.WriteFile(file, []byte("File1\nLine 2\nEnd of file"), 0644)
	gist2, _ := CreateGistFromLocalFiles("Test 2", file)

	wantDiff := []string{
		"file",
		"  \x1b[31mf\x1b[0m\x1b[32mF\x1b[0mile1\n\x1b[31ml\x1b[0m\x1b[32mL\x1b[0mine 2\n\x1b[31me\x1b[0m\x1b[32mE\x1b[0mnd of file",
	}

	t.Run("Default", func(t *testing.T) {
		gotDiff := getFilesDiff(gist1, gist2)
		if !reflect.DeepEqual(gotDiff, wantDiff) {
			t.Errorf("getFilesDiff() = %v, want %v", gotDiff, wantDiff)
		}
	})

}
