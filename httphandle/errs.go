package httphandle

import "net/http"

type ErrorResponse struct {
	Code    int
	Message string
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func BadRequest(msg string) error{
	return ErrorResponse{
		Code: http.StatusBadRequest,
		Message: msg,
	}
}

func NotFound(msg string) error{
	return ErrorResponse{
		Code: http.StatusNotFound,
		Message: msg,
	}
}
