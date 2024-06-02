package routeerrors

import (
	"backend/app/routes/authenticator/cookie"
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

func Forbidden() *RequestError {
	return &RequestError{
		StatusCode: http.StatusForbidden,
		Message:    "Forbidden",
	}
}

func NotFound(message string) *RequestError {
	return &RequestError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func HandleError(w http.ResponseWriter, err error) {
	var re *RequestError
	if errors.As(err, &re) {
		if re.StatusCode == http.StatusUnauthorized || re.StatusCode == http.StatusForbidden {
			cookie.DeleteCookie(w)
		}

		http.Error(w, re.Message, re.StatusCode)
		return
	}

	http.Error(w, err.Error(), http.StatusInternalServerError)
}
