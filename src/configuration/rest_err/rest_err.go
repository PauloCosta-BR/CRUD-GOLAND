package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}
type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message string, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "Bad Request",
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "Bad Request",
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     "Internal Server Error",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     "Resource Not Found",
	}
}

func NewMethodNotAllowedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusMethodNotAllowed,
		Err:     "Method Not Allowed",
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusUnauthorized,
		Err:     "Unauthorized",
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusForbidden,
		Err:     "Forbidden",
	}
}

func NewConflictError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusConflict,
		Err:     "Conflict",
	}
}

func NewUnprocessableEntityError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
		Err:     "Unprocessable Entity",
	}
}

func (restErr *RestErr) Error() string {
	return restErr.Message
}

func (restErr *RestErr) StatusCode() int {
	return restErr.Code
}

func (restErr *RestErr) Cause() string {
	return restErr.Causes[0].Field
}
