package status

import "errors"

// Status handling errors
type Status struct {
	StatusCode int
	Error error
}

// NewStatus sends pointer to status message
func NewStatus(statusCode int, err error) *Status {
	return &Status{StatusCode: statusCode, Error: err}
}

// Ok returns successful message
func Ok() *Status {
	return NewStatus(200, nil)
}

// NotFound error
func NotFound() *Status {
	err := errors.New("Not Found")
	return &Status{StatusCode: 400, Error: err}
}
