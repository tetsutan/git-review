package review

import (
	"github.com/tetsutan/git-review/src/review/model"
	"github.com/tetsutan/git-review/src/review/common"
	"fmt"
	"os"
)

type Review struct {

}


func (r *Review) Start(spec string) bool {

	// currentDirectory, _ := filepath.Abs(".")

	wrapper := model.GitWrapper{}

	fmt.Println("Reviewing start")
	fmt.Println("rootPath =>", wrapper.RootPathByGit())

	state := model.StateRepository{
		RootPath: wrapper.RootPathByGit(),
	}

	if !state.GitAvailable() {
		fmt.Println(state.RootPath, " is not git repository")
		return false
	}

	if !wrapper.IsAvailableSpec(state.RootPath, spec) {
		fmt.Println("Spec (", spec, ") couldn't  be found")
		return false
	}

	// TODO 現在参照しているディレクトリ以下岳を見るフラグを追加
	commitIds := wrapper.GetCommitIdsFromSpec(spec)

	fmt.Println(len(commitIds), "commits to review")

	state.Init(spec, commitIds)

	current := state.GetCurrentCommitId()
	state.WriteLog("[Start] " + current)

	fmt.Println("")
	fmt.Println("Current commit is", current)

	fmt.Println("--- Commit message ---")
	fmt.Println(wrapper.GetCommitMessage(current))
	body := wrapper.GetCommitMessageBody(current)
	if len(body) > 0 {
		fmt.Println(body)
	}

	return true
}

func (r *Review) Reset() bool {

	wrapper := model.GitWrapper{}
	state := model.StateRepository{
		RootPath: wrapper.RootPathByGit(),
	}

	if !state.IsReviewing() {
		fmt.Println("Not reviewing")
		return false
	}

	state.ClearAllFiles()
	return true

}

func (r *Review) next(uncommentedOnly bool) bool {
	Log("Next")
	wrapper := model.GitWrapper{}
	state := model.StateRepository{
		RootPath: wrapper.RootPathByGit(),
	}

	if !state.IsReviewing() {
		fmt.Println("Not reviewing")
		return false
	}

	commitId, err := state.GetNextCommitId(uncommentedOnly)
	Log("NexCommitID "+ commitId)

	if err != nil {
		fmt.Println("err", err)
	}

	if len(commitId) > 0 {
		Log("Next Call WriteCurrent "+ commitId)
		state.WriteCurrent(commitId)
		state.WriteLog("[Next] switch to " + commitId)
		r.Status()
	} else if len(state.GetCurrentCommitId()) > 0 {
		fmt.Println("Previous is last commit to review.")


		// calc not commented
		commitIds := state.GetAllCommitIds()
		commitSet := &common.StringSet{}
		commitSet.SetAll(commitIds)

		commentedIds := state.GetCommentedIds()
		commentedSet := &common.StringSet{}
		commentedSet.SetAll(commentedIds)
		notCommentedSet := commitSet.Minus(commentedSet)
		notCommentedIds := notCommentedSet.All()
		notCommentedSize := len(notCommentedIds)

		if notCommentedSize > 0 {
			fmt.Println("But not commented is ", notCommentedSize, "commits below.")
			for _, id := range notCommentedIds {
				fmt.Println(id)
			}
		}

		fmt.Println("Use `git review complete` if review is finished? ")
		// r.Reset()
	}

	return true
}
func (r *Review) Next() bool {
	return r.next(false)
}
func (r *Review) NextUncommented() bool {
	return r.next(true)
}

func (r *Review) prev(uncommentedOnly bool) bool {
	Log("Prev")
	wrapper := model.GitWrapper{}
	state := model.StateRepository{
		RootPath: wrapper.RootPathByGit(),
	}

	if !state.IsReviewing() {
		fmt.Println("Not reviewing")
		return false
	}

	commitId, err := state.GetPrevCommitId(uncommentedOnly)
	Log("PrevCommitID "+ commitId)

	if err != nil {
		fmt.Println("err", err)
	}

	if len(commitId) > 0 {
		Log("Prev Call WriteCurrent "+ commitId)
		state.WriteCurrent(commitId)
		state.WriteLog("[Prev] switch to " + commitId)
		r.Status()
	}

	return true
}
func (r *Review) Prev() bool {
	return r.prev(false)
}
func (r *Review) PrevUncommented() bool {
	return r.prev(true)
}

func (r *Review) Status() bool {

	wrapper := model.GitWrapper{}
	state := model.StateRepository{
		RootPath: wrapper.RootPathByGit(),
	}

	if !state.IsReviewing() {
		fmt.Fprintln(os.Stderr, "Not reviewing")
		return false
	}

	current := state.GetCurrentCommitId()

	commitIds := state.GetAllCommitIds()
	commitsCount := len(commitIds)
	fmt.Print(commitsCount, " commits to review")

	for i, id := range commitIds {
		if id == current {
			fmt.Print(", ", commitsCount - i, " remain")
		}
	}

	commitSet := &common.StringSet{}
	commitSet.SetAll(commitIds)

	commentedIds := state.GetCommentedIds()
	commentedSet := &common.StringSet{}
	commentedSet.SetAll(commentedIds)
	notCommentedSet := commitSet.Minus(commentedSet)
	notCommentedSize := len(notCommentedSet.All())
	fmt.Print(", ", notCommentedSize, " not-commented")

	fmt.Println("")

	fmt.Println("")
	fmt.Println("Current commit is", current)

	fmt.Println("--- Commit message ---")
	fmt.Println(wrapper.GetCommitMessage(current))
	body := wrapper.GetCommitMessageBody(current)
	if len(body) > 0 {
		fmt.Println(body)
	}

	return true
}

func (r *Review) Comments() bool {

	wrapper := model.GitWrapper{}
	state := model.StateRepository{
		RootPath: wrapper.RootPathByGit(),
	}

	if !state.IsReviewing() {
		fmt.Println("Not reviewing")
		return false
	}

	commitIds := state.GetAllCommitIds()
	comments := state.GetComments()

	fmt.Println("Show comments")
	for _, id := range commitIds {

		fmt.Print("# id: ", id)
		fmt.Println("")
		if len(comments[id]) > 0 {
			fmt.Print(comments[id])
			fmt.Println("")
		}
		fmt.Println("")

	}

	return true
}


func (r *Review) Comment(message string) bool {
	Log("Comment")
	return r.comment(common.Message, message)
}

func (r *Review) comment(commentType common.CommentType, message string) bool {

	wrapper := model.GitWrapper{}
	state := model.StateRepository{
		RootPath: wrapper.RootPathByGit(),
	}

	if !state.IsReviewing() {
		fmt.Println("Not reviewing")
		return false
	}

	current := state.GetCurrentCommitId()

	if len(current) > 0 {
		state.WriteComment(current, commentType, message)
		state.WriteLog("[Comment] comment to " + current)
	}

	return true
}

func (r *Review) Good() bool {
	Log("Good")
	if !r.comment(common.Good, "") {
		return false
	}
	return r.Next()
}

func (r *Review) Bad() bool {
	Log("Bad")
	if !r.comment(common.Bad, "") {
		return false
	}
	return r.Next()
}

func (r *Review) Skip() bool {
	Log("Skip")
	if !r.comment(common.Skip, "") {
		return false
	}
	return r.Next()
}

func (r *Review) Complete() bool {

	wrapper := model.GitWrapper{}
	state := model.StateRepository{
		RootPath: wrapper.RootPathByGit(),
	}

	if !state.IsReviewing() {
		fmt.Println("Not reviewing")
		return false
	}

	completeFilePath := state.WriteComplete()
	r.Reset()

	fmt.Println("Review comments are written to " + completeFilePath)

	return true
}
