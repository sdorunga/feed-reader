package app

import (
	"net/http"
)

// HttpError is an interface that the router uses to decide which status code
// it should apply to the errors it receives from the handlers
type HttpError interface {
	StatusCode() int
	error
}

// NotFoundError wraps an error with status code 404
type NotFoundError struct {
	err error
}

func (notFound NotFoundError) Error() string {
	return notFound.err.Error()
}

func (notFound NotFoundError) StatusCode() int {
	return http.StatusNotFound
}

// BadRequestError wraps an error with status code 400
type BadRequestError struct {
	err error
}

func (badRequest BadRequestError) Error() string {
	return badRequest.err.Error()
}

func (notFound BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}
