package model

import (
	"os"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"bufio"
	"strings"
	"github.com/pkg/errors"
	"strconv"
	"github.com/tetsutan/git-review/src/review/common"
	"time"
)

type StateRepository struct {
	RootPath string
}

var ReviewCurrent = "REVIEW_CURRENT"
var ReviewAll = "REVIEW_ALL"
var ReviewLog = "REVIEW_LOG"
var ReviewComment = "REVIEW_COMMENT"

var ReviewFiles = []string{
	ReviewCurrent,
	ReviewAll,
	ReviewLog,
	ReviewComment,
}

func (state *StateRepository) Init(spec string, commitIds []string) {
	state.ClearAllFiles()
	state.WriteLog("# Initial spec is " + spec)
	state.WriteAll(commitIds)
	state.WriteCurrent(commitIds[0])
	state.TouchCommentFile()

}

func (state *StateRepository)ClearAllFiles() {

	if state.GitAvailable() {
		for _, name := range ReviewFiles {
			filePath := filepath.Join(state.GetDotGitPath(), name)
			if err := os.Remove(filePath); err != nil {
				// fmt.Println(err)
			}
		}

	}

}

func (state *StateRepository) GitAvailable() bool {
	_, err := os.Stat(state.GetDotGitPath())
	return err == nil
}

func (state *StateRepository) GetDotGitPath() string {
	return filepath.Join(state.RootPath, ".git")
}

func (state *StateRepository) IsReviewing() bool {
	_,err := os.Stat(filepath.Join(state.GetDotGitPath(), ReviewCurrent))
	return err == nil
}

func (state *StateRepository) WriteLog(s string) {
	if state.GitAvailable() {
		file, err := os.OpenFile(filepath.Join(state.GetDotGitPath(), ReviewLog),
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		file.WriteString(s)
		file.WriteString("\n")
	}
}
func (state *StateRepository) WriteCurrent(s string) {
	if state.GitAvailable() {
		file, err := os.OpenFile(filepath.Join(state.GetDotGitPath(), ReviewCurrent),
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		file.WriteString(s)
	}
}
func (state *StateRepository) WriteAll(ss []string) {
	if state.GitAvailable() {
		file, err := os.OpenFile(filepath.Join(state.GetDotGitPath(), ReviewAll),
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		for _, s := range ss {
			file.WriteString(s)
			file.WriteString("\n")
		}

	}
}
func (state *StateRepository) WriteComment(commitId string, commentType common.CommentType, message string) {
	if state.GitAvailable() {
		file, err := os.OpenFile(filepath.Join(state.GetDotGitPath(), ReviewComment),
			os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		file.WriteString(commitId)
		file.WriteString("\t")
		file.WriteString(fmt.Sprintf("%d", commentType))
		file.WriteString("\t")
		file.WriteString(strings.Replace(message, "\r", " ", -1))
		file.WriteString("\n")
	}
}
func (state *StateRepository) TouchCommentFile() {
	if state.GitAvailable() {
		file, err := os.Create(filepath.Join(state.GetDotGitPath(), ReviewComment))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}
}

// Writing Path is RootPath
func (state *StateRepository) WriteComplete() (path string) {
	commitIds := state.GetAllCommitIds()
	comments := state.GetComments()

	if len(commitIds) > 0 {

		nowTime := time.Now()
		path = "review-complete-" + nowTime.Format("20060102-150405" )+".txt"
		file, err := os.OpenFile(filepath.Join(state.RootPath, path), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		file.WriteString("# " + nowTime.Format("2006/01/02-15:04:05" ) + "\n")
		for _, id := range commitIds {
			file.WriteString("id: " + id + "\n")
			file.WriteString(comments[id])
			file.WriteString("\n\n")
		}

	}

	return

}

func (state *StateRepository) GetCurrentCommitId() string {
	if state.GitAvailable() {
		data, err := ioutil.ReadFile(filepath.Join(state.GetDotGitPath(), ReviewCurrent))
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return string(data)
	}
	return ""
}

func (state *StateRepository) GetNextCommitId(uncommentedOnly bool) (string, error) {
	current := state.GetCurrentCommitId()

	if len(current) > 0 {

		f, err := os.Open(filepath.Join(state.GetDotGitPath(), ReviewAll))
		if err != nil {
			fmt.Println(err)
			return "", errors.New("Can not open " + ReviewAll)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		found := false

		commentedSet := &common.StringSet{}
		if uncommentedOnly {
			commentedSet.SetAll(state.GetCommentedIds())
		}

		for scanner.Scan() {
			// appendで追加
			text := strings.TrimSpace(scanner.Text())
			if found {
				if uncommentedOnly {
					if !commentedSet.Contains(text) {
						return text, nil
					}
				} else {
					return text, nil
				}
			}
			if text == current {
				found = true

			}
		}

		if found {
			return "", nil // last

		}
	}
	return "", errors.New("Invalid commit ID")
}

func (state *StateRepository) GetPrevCommitId(uncommentedOnly bool) (string,error) {
	current := state.GetCurrentCommitId()
	if len(current) > 0 {

		f, err := os.Open(filepath.Join(state.GetDotGitPath(), ReviewAll))
		if err != nil {
			fmt.Println(err)
			return "", errors.New("Can not open " + ReviewAll)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		prevId := ""

		commentedSet := &common.StringSet{}
		if uncommentedOnly {
			commentedSet.SetAll(state.GetCommentedIds())
		}

		for scanner.Scan() {
			// appendで追加
			text := strings.TrimSpace(scanner.Text())
			if text == current {
				return prevId, nil
			}

			if uncommentedOnly {
				if !commentedSet.Contains(text) {
					prevId = text
				}
			} else {
				prevId = text
			}
		}

		if len(prevId) == 0 {
			return "", errors.New(ReviewAll + " does not contain IDs")
		} else {
			return "", errors.New(current + " is not included in " + ReviewAll)
		}
	}
	return "", errors.New("Invalid commit ID")
}


func (state *StateRepository) GetAllCommitIds() (commitIds []string) {
	if state.GitAvailable() {
		data, err := ioutil.ReadFile(filepath.Join(state.GetDotGitPath(), ReviewAll))
		if err != nil {
			fmt.Println(err)
			return
		}

		commitIds = strings.Split(strings.TrimSpace(string(data)), "\n")
	}

	return
}

func (state *StateRepository) Reviewing() bool {
	return len(state.GetCurrentCommitId()) > 0
}

func (state *StateRepository) GetCommentedIds() (commitIds []string) {
	if state.GitAvailable() {
		data, err := ioutil.ReadFile(filepath.Join(state.GetDotGitPath(), ReviewComment))
		if err != nil {
			fmt.Println(err)
			return
		}

		commitAndMessage := strings.Split(strings.TrimSpace(string(data)), "\n")

		set := common.StringSet{}
		for _, line := range commitAndMessage {
			splits := strings.SplitN(line, "\t", 3)
			if len(splits) >= 2 {
				commitId := strings.TrimSpace(splits[0])
				commentType,_ := strconv.Atoi(splits[1])

				if commentType != int(common.Skip) {
					set.Set(commitId)
				}

			}
		}

		commitIds = set.All()
	}

	return
}



func (state *StateRepository) GetComments() (comments map[string]string) {

	comments = make(map[string]string)

	if state.GitAvailable() {
		data, err := ioutil.ReadFile(filepath.Join(state.GetDotGitPath(), ReviewComment))
		if err != nil {
			fmt.Println(err)
			return
		}

		commitAndMessage := strings.Split(strings.TrimSpace(string(data)), "\n")

		commentTypes := make(map[string]common.CommentType)
		commentMessages := make(map[string][]string)

		// typeは後ろにあるので上書きされる
		// メッセージは追記
		for _, line := range commitAndMessage {
			splits := strings.SplitN(line, "\t", 3)
			if len(splits) >= 2 {
				commitId := strings.TrimSpace(splits[0])
				commentType,_ := strconv.Atoi(splits[1])
				commentTypes[commitId] = common.CommentType(commentType)

				if len(splits) >= 3 {
					commitMessage := strings.TrimSpace(splits[2])
					commentMessages[commitId] = append(commentMessages[commitId], commitMessage)
				}
			}
		}

		for id, commentType := range commentTypes {
			comments[id] = "*" + common.CommentTypeToString(commentType) + "*\n" +
				strings.Join(commentMessages[id], "\n")
		}
	}

	return

}
