package util

import "net/http"

// Error ...
type Error struct {
	Status int    // http code
	Errno  int    // biz error code
	ErrMsg string // biz error message
}

var (
	ErrSuccess        = &Error{http.StatusOK, 0, "success"}
	ErrInvalidParam   = &Error{http.StatusOK, 40000, "invalid param"}
	ErrInternalServer = &Error{http.StatusOK, 50000, "internal server error"}
)

// Error implement error interface
func (e *Error) Error() string {
	return e.ErrMsg
}

// More ...
func (e *Error) More(msg string) *Error {
	return &Error{e.Status, e.Errno, e.ErrMsg + ": " + msg}
}

// CastError ...
func CastError(err error) *Error {
	return CastErrorDefault(err, ErrInternalServer)
}

// CastErrorDefault ...
func CastErrorDefault(err error, defaultErr *Error) *Error {
	if err == nil {
		return ErrSuccess
	}
	if bizErr, ok := err.(*Error); ok {
		return bizErr
	}

	return defaultErr.More(err.Error())
}
