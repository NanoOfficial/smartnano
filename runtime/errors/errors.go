package runtimeerrors

import (
	"fmt"
	"runtime/debug"

	"golang.org/x/xerrors"
)

type InternalError interface {
	error
	IsInternalError()
}

type UserError interface {
	error
	IsUserError()
}

type ExternalError struct {
	Recovered error
}

type ExternalNonError struct {
	Recovered any
}

type SecondaryError interface {
	SecondaryError() string
}

type ErrorNotes interface {
	ErrorNotes() []ErrorNote
}

type ErrorNote interface {
	Message() string
}

type ParentError interface {
	error
	ChildErrors() []error
}

type HasPrefix interface {
	Prefix() string
}

type MemoryError struct {
	Err error
}

type UnexpectedError struct {
	Err   error
	Stack []byte
}

type DefaultUserError struct {
	Err error
}

func NewUnreachableError() InternalError {
	return NewUnexpectedError("unreachable")
}

func NewExternalError(recovered error) ExternalError {
	return ExternalError{
		Recovered: recovered,
	}
}

func (e ExternalError) Error() string {
	return fmt.Sprint(e.Recovered)
}

func (e ExternalError) Unwrap() error {
	return e.Recovered
}

func NewExternalNonError(recovered error) ExternalError {
	return ExternalError{
		Recovered: recovered,
	}
}

func (e ExternalNonError) Error() string {
	return fmt.Sprint(e.Recovered)
}

var _ UserError = MemoryError{}

func (MemoryError) IsUserError() {}

func (e MemoryError) Unwrap() error {
	return e.Err
}

func (e MemoryError) Error() string {
	return fmt.Sprintf("memory error: %s", e.Err.Error())
}

var _ InternalError = UnexpectedError{}

func (UnexpectedError) IsInternalError() {}

func NewUnexpectedError(message string, arg ...any) UnexpectedError {
	return UnexpectedError{
		Err:   fmt.Errorf(message, arg...),
		Stack: debug.Stack(),
	}
}

func NewUnexpectedErrorFromCause(err error) UnexpectedError {
	return UnexpectedError{
		Err:   err,
		Stack: debug.Stack(),
	}
}

func (e UnexpectedError) Unwrap() error {
	return e.Err
}

func (e UnexpectedError) Error() string {
	return fmt.Sprintf("internal error: %s\n%s", e.Err.Error(), e.Stack)
}

var _ UserError = DefaultUserError{}

func (DefaultUserError) IsUserError() {}

func NewDefaultUserError(message string, arg ...any) DefaultUserError {
	return DefaultUserError{
		Err: fmt.Errorf(message, arg...),
	}
}

func (e DefaultUserError) Unwrap() error {
	return e.Err
}

func (e DefaultUserError) Error() string {
	return e.Err.Error()
}

func IsInternalError(err error) bool {
	switch err := err.(type) {
	case InternalError:
		return true
	case xerrors.Wrapper:
		return IsInternalError(err.Unwrap())
	default:
		return false
	}
}

func IsUserError(err error) bool {
	switch err := err.(type) {
	case UserError:
		return true
	case xerrors.Wrapper:
		return IsUserError(err.Unwrap())
	default:
		return false
	}
}

func GetExternalError(err error) (ExternalError, bool) {
	switch err := err.(type) {
	case ExternalError:
		return err, true
	case xerrors.Wrapper:
		return GetExternalError(err.Unwrap())
	default:
		return ExternalError{}, false
	}
}
