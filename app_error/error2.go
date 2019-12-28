package app_error

import (
	"errors"
	"fmt"
)

type badRequest interface {
	BadRequest() bool
}
type unauthorized interface {
	Unauthorized() bool
}

type notFound interface {
	NotFound() bool
}

type internal interface {
	Internal() bool
}

func IsBadRequest(err error) bool {
	var br badRequest
	return errors.As(err, &br)
}

func IsUnauthorized(err error) bool {
	var un unauthorized
	return errors.As(err, &un)
}

func IsNotFound(err error) bool {
	var nf notFound
	return errors.As(err, &nf)
}

type ValidationErr struct{ s string }
type DuplicatedIdErr struct{ s string }
type UnauthorizedErr struct{ s string }
type ResourceNotFoundErr struct{ s string }
type UnknownErr struct{ s string }

func (e ValidationErr) Error() string    { return e.s }
func (e ValidationErr) BadRequest() bool { return true }

func (e DuplicatedIdErr) Error() string    { return e.s }
func (e DuplicatedIdErr) BadRequest() bool { return true }

func (e UnauthorizedErr) Error() string      { return e.s }
func (e UnauthorizedErr) Unauthorized() bool { return true }

func (e ResourceNotFoundErr) Error() string  { return e.s }
func (e ResourceNotFoundErr) NotFound() bool { return true }

func NewAppErr2(typeS string) error {
	if typeS == "ValidationErr" {
		return ValidationErr{"parameter is invalid"}
	}
	if typeS == "WrappedValidationErr" {
		return fmt.Errorf("error wrapped: %w", ValidationErr{"parameter is invalid"})
	}
	if typeS == "DuplicatedIdErr" {
		return DuplicatedIdErr{"id is duplicated"}
	}
	if typeS == "UnauthorizedErr" {
		return UnauthorizedErr{"unauthorized"}
	}
	if typeS == "ResourceNotFoundErr" {
		return ResourceNotFoundErr{"resource not found"}
	}
	return errors.New("not defined error occurred")
}
