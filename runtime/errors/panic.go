package runtimeerrors

import (
	goRuntime "runtime"
)

func WrapPanic(f func()) {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case goRuntime.Error, InternalError:
				panic(r)
			case error:
				panic(ExternalError{
					Recovered: r,
				})
			default:
				panic(ExternalNonError{
					Recovered: r,
				})
			}
		}
	}()
	f()
}
