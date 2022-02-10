package error

import "net/http"

type HttpErr interface {
	Error() string
	Code() int
}

/*
InternalServerErr

Status:500

some unexpected error occurred by server side.
*/
type InternalServerErr struct {
	Origin error
}

func (i InternalServerErr) Error() string {
	return i.Origin.Error()
}
func (i InternalServerErr) Code() int {
	return http.StatusInternalServerError
}

/*
Forbidden

Status:403

client hasn't access rights.
*/
type Forbidden struct {
	Origin error
}

func (f Forbidden) Error() string {
	return f.Origin.Error()
}
func (f Forbidden) Code() int {
	return http.StatusForbidden
}

/*
UnAuthorized

Status:401

invalid param to authorize.
*/
type UnAuthorized struct {
	Origin error
}

func (u UnAuthorized) Error() string {
	return u.Origin.Error()
}
func (u UnAuthorized) Code() int {
	return http.StatusUnauthorized
}
