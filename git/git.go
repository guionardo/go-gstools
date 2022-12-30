package git

import (
	"os/exec"
	"strings"
)

type GitUser struct {
	Name  string
	Email string
}

func GetCurrentGitUser() (gitUser *GitUser, err error) {
	var out []byte
	out, err = exec.Command("git", "config", "--list").Output()
	if err != nil {
		return
	}
	gitUser = &GitUser{}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		words := strings.Split(line, "=")
		if len(words) != 2 {
			continue
		}
		switch words[0] {
		case "user.name":
			gitUser.Name = words[1]
		case "user.email":
			gitUser.Email = words[1]
		}
	}
	return gitUser, nil
}
