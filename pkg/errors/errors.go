package util_error

import "fmt"

type ErrorLevel int

const (
	Repository ErrorLevel = iota + 1
	UseCase
	Driver
)

type ErrorStruct struct {
	Level   ErrorLevel
	Message string
}

func GetErrorLevel(err ErrorStruct) ErrorLevel {
	return err.Level
}

func GetErrorMessage(err ErrorStruct) string {
	return err.Message
}

func NewError(err error, l ErrorLevel, msg ...string) *ErrorStruct {
	if len(msg) > 0 {
		return &ErrorStruct{
			Level:   l,
			Message: fmt.Errorf("%s: %w", msg[0], err).Error(),
		}
	}
	return &ErrorStruct{
		Level:   l,
		Message: err.Error(),
	}
}
func NewRepositoryError(err error, msg ...string) *ErrorStruct {
	if len(msg) > 0 {
		return &ErrorStruct{
			Level:   Repository,
			Message: fmt.Errorf("%s: %w", msg[0], err).Error(),
		}
	}
	return &ErrorStruct{
		Level:   Repository,
		Message: err.Error(),
	}
}
func NewUseCaseError(err error, msg ...string) *ErrorStruct {
	if len(msg) > 0 {
		return &ErrorStruct{
			Level:   UseCase,
			Message: fmt.Errorf("%s: %w", msg[0], err).Error(),
		}
	}
	return &ErrorStruct{
		Level:   UseCase,
		Message: err.Error(),
	}
}

func (e *ErrorStruct) Error() string {
	return e.Message
}
