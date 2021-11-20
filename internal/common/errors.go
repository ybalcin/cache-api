package common

type (
	// Error interface
	Error interface {
		error
		Status() int
	}

	// StatusError struct is status error
	StatusError struct {
		Code int
		Err  error
	}
)

// NewStatusError initializes new status err
func NewStatusError(code int, err error) error {
	return StatusError{
		Code: code,
		Err:  err,
	}
}

// Status returns http status code
func (e StatusError) Status() int {
	return e.Code
}

// Error returns error message
func (e StatusError) Error() string {
	return e.Err.Error()
}
