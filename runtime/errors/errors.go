package errors

type InternalError interface {
	error
	IsInternalError()
}

type ExternalError struct {
	Recovered error
}

type UserError interface {
	error
	isUserError()
}

type ExternalNonError struct {
	Recovered any
}
