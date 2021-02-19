package strava

const (
	OpErrorAccessDenied OpErrorType = iota
	OpErrorCodeInvalid
	OpErrorClientIDInvalid
	OpErrorClientSecretInvalid
	OpErrorGrantTypeInvalid
)

type OpErrorType int

type OpError struct {
	err     error
	errtype OpErrorType
	code    int
}

func (e OpError) Error() string {
	return e.err.Error()
}
