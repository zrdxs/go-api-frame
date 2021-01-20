package error

import (
	"fmt"
	"runtime"
	"strings"
)

// WrapError interface
type WrapError interface {
	Error() string
	OriginalError() error
	Message() error
}

// WrappedError struct
type WrappedError struct {
	originalError error
	path          string
	messages      []string
}

func (err *WrappedError) Error() string {
	if len(err.messages) > 0 {
		pathErr := fmt.Sprintf("%s", err.path)

		for _, message := range err.messages {
			pathErr += fmt.Sprintf("%s; ", message)
		}

		return fmt.Sprintf("%s -- %v", pathErr, err.originalError)
	}

	return fmt.Sprintf("%s -- %v", err.path, err.originalError)
}

// GetOriginalError returns the original error
func (err *WrappedError) GetOriginalError() error {
	if err.originalError != nil {
		wrappedErr, ok := (err.originalError).(WrapError)
		if ok {
			return wrappedErr.OriginalError()
		}
	}

	return err.originalError
}

// Wrap returns the error Wrapped
func Wrap(err error, messages ...string) error {
	//Get function path
	pc := make([]uintptr, 1)
	runtime.Callers(4, pc)

	funcRef := runtime.FuncForPC(pc[0])
	funcPath := strings.Split(funcRef.Name(), "/")

	if err != nil {
		return &WrappedError{
			originalError: err,
			path:          funcPath[len(funcPath)-1],
			messages:      messages,
		}
	}

	return nil
}
