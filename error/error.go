package error

import "fmt"

type AppError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("error:  %s", e.Message)
}

func BadRequestError(m string) error {
	return newAppError(400, m)
}
func ConflictError(m string) error {
	return newAppError(409, m)
}

func newAppError(statusCode int, message string) AppError {
	return AppError{
		StatusCode: statusCode,
		Message:    message,
	}
}
