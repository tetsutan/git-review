package common

type CommentType int
const (
	Message CommentType = iota
	Good
	Bad
	Skip
)

func CommentTypeToString(commentType CommentType) string {
	switch commentType {
	case Message: return "Message"
	case Good: return "Good"
	case Bad: return "Bad"
	case Skip: return "Skip"
	}

	return ""
}

