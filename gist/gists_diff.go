package gist

import (
	"github.com/google/go-github/v48/github"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func getFilesDiff(gist1 *github.Gist, gist2 *github.Gist) (diff []string) {
	diffs := make(map[string][]string)
	assertDiffs := func(filename string) {
		if _, ok := diffs[filename]; !ok {
			diffs[filename] = make([]string, 0, 20)
		}
	}
	dmp := diffmatchpatch.New()

	if gist1 != nil && len(gist1.Files) > 0 {
		for fileName, file := range gist1.Files {
			assertDiffs(*file.Filename)

			if gist2 == nil || len(gist2.Files) == 0 {
				diffs[*file.Filename] = append(diffs[*file.Filename], "File not found in remote")
			} else if file2, ok := gist2.Files[fileName]; !ok {
				diffs[*file.Filename] = append(diffs[*file.Filename], "File not found in remote")
			} else {
				d := dmp.DiffMain(*file.Content, *file2.Content, true)
				diffs[*file.Filename] = append(diffs[*file.Filename], dmp.DiffPrettyText(d))
			}
		}
	}

	if gist2 != nil && len(gist2.Files) > 0 {
		for fileName, file := range gist2.Files {
			if _, ok := diffs[*file.Filename]; ok {
				continue
			}
			assertDiffs(*file.Filename)

			if gist1 == nil || len(gist1.Files) == 0 {
				diffs[*file.Filename] = append(diffs[*file.Filename], "File not found in local")
			} else if file2, ok := gist1.Files[fileName]; !ok {
				diffs[*file.Filename] = append(diffs[*file.Filename], "File not found in local")
			} else {
				d := dmp.DiffMain(*file2.Content, *file.Content, true)
				diffs[*file.Filename] = append(diffs[*file.Filename], dmp.DiffPrettyText(d))
			}
		}
	}
	for fileName, d := range diffs {
		diff = append(diff, fileName)
		for _, line := range d {
			diff = append(diff, "  "+line)
		}
	}

	return
}
