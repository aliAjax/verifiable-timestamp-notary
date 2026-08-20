package protocol

type Code string

const (
	InvalidRequest Code = "invalid-request"
	BadDigest      Code = "bad-digest"
	Unauthorized   Code = "unauthorized"
	QuorumFailure  Code = "quorum-failure"
	Internal       Code = "internal"
)

type Error struct {
	Code      Code
	Message   string
	Retryable bool
	Detail    map[string]string
}

func (e Error) Error() string    { return string(e.Code) + ": " + e.Message }
func New(c Code, m string) Error { return Error{Code: c, Message: m, Detail: map[string]string{}} }
func (e Error) With(k, v string) Error {
	if e.Detail == nil {
		e.Detail = map[string]string{}
	}
	e.Detail[k] = v
	return e
}
