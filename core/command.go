package core

const (
	CommandCodeReviewRequest = iota + 1
)

var commandText = map[int]string{
	CommandCodeReviewRequest: "code-review request",
}

func CommandText(code int) string {
	return commandText[code]
}
