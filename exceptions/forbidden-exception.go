package exceptions

type forbiddenError struct {
	errorMsg string
}

func NewForbiddenError(msg string) *forbiddenError {
	return &forbiddenError{
		errorMsg: msg,
	}
}

func (e *forbiddenError) Error()
