package api

type Problem struct {
	Type    string
	Title   string
	Status  int
	Detail  string
	TraceID string
}

func NewProblem(status int, title, detail string) Problem {
	return Problem{Type: "about:blank", Title: title, Status: status, Detail: detail}
}
