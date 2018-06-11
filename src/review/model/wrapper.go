package model

import (
	"os/exec"
	"fmt"
	"strings"
)

type GitWrapper struct{}

func (gw *GitWrapper) IsAvailableSpec(path string, spec string) bool {
	cmd := exec.Command("git", "rev-parse", spec)
	_, err := cmd.Output()
	return err == nil
}

func (gw *GitWrapper) RootPathByGit() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	return strings.TrimSpace(wrapCommand(cmd.Output()))
}

func wrapCommand(out []byte, err error) string{
	if err != nil {
		fmt.Println("cmd err", err)
		return ""
	}
	return string(out)
}

func (gw *GitWrapper) GetCommitIdsFromSpec(spec string) (commitIds []string) {

	// git log --no-merges -s --format='%H' $review_commit 2>/dev/null | tail -r
	// cmd := exec.Command("git", "rev-parse", "--no-merges", spec)
	cmd := exec.Command("git", "log", "--no-merges", "-s", "--format=%H", spec)
	lineStrings := wrapCommand(cmd.Output())

	lines := strings.Split(lineStrings, "\n")

	for _, line := range lines {
		// unshift
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			commitIds = append([]string{ line }, commitIds...)
		}

	}

	return
}


func (gw *GitWrapper) GetCommitMessage(commitId string) string {
	// git show --no-merges -s --format='%b'
	cmd := exec.Command("git", "show", "--no-merges", "-s", "--format=%s", commitId)
	return strings.TrimSpace(wrapCommand(cmd.Output()))
}

func (gw *GitWrapper) GetCommitMessageBody(commitId string) string {
	// git show --no-merges -s --format='%b'
	cmd := exec.Command("git", "show", "--no-merges", "-s", "--format=%b", commitId)
	return strings.TrimSpace(wrapCommand(cmd.Output()))
}
