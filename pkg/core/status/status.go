package status

import (
	"fmt"
	"net/http"
)

// ErrInternal ...
var ErrInternal = ErrServiceStatus{
	ServiceStatus{Code: http.StatusInternalServerError, Message: "Internal Server Error"},
}

// ErrNotFound ...
var ErrNotFound = ErrServiceStatus{
	ServiceStatus{Code: http.StatusNotFound, Message: "Not Found"},
}

// ErrBadRequest ...
var ErrBadRequest = ErrServiceStatus{
	ServiceStatus{Code: http.StatusBadRequest, Message: "Bad Request"},
}

// ErrUnauhtorized ...
var ErrUnauhtorized = ErrServiceStatus{
	ServiceStatus{Code: http.StatusUnauthorized, Message: "Unauthorized"},
}

// Success ...
var Success = ServiceStatus{Code: http.StatusOK, Message: "OK"}

// ServiceStatus ...
type ServiceStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrServiceStatus ...
type ErrServiceStatus struct {
	ServiceStatus
}

func (e ErrServiceStatus) Error() string {
	return fmt.Sprintf(string(e.Code), ": ", e.Message)
}

// WithMessage ...
func (e ErrServiceStatus) WithMessage(msg string) ErrServiceStatus {
	return ErrServiceStatus{ServiceStatus{Code: e.Code, Message: msg}}
}

// New ...
func New(ss ServiceStatus) ServiceStatus {
	return ServiceStatus{ss.Code, ss.Message}
}

// NewUserDefined ...
func NewUserDefined(code int, msg string) ServiceStatus {
	return ServiceStatus{Code: code, Message: msg}
}
