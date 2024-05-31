package routeerrors

import (
	"errors"
	"fmt"
	"net/http"
)

type RequestError struct {
	StatusCode int
	Message    string
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("status code: %d, message: %s", e.StatusCode, e.Message)
}

func BadRequest(message string) *RequestError {
	return &RequestError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func Unauthorized() *RequestError {
	return &RequestError{
		StatusCode: http.StatusUnauthorized,
		Message:    "Not authorized",
	}
}

func HandleError(w http.ResponseWriter, err error) {
	var re *RequestError
	if errors.As(err, &re) {
		http.Error(w, re.Message, re.StatusCode)
		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}
