package result

import "fmt"

type Result[T any] struct {
	value T
	err   error
}

func Ok[T any](val T) Result[T] {
	return Result[T]{
		value: val,
		err:   nil,
	}
}

func Error[T any](err error) (r Result[T]) {
	r.err = err
	return
}

func (r Result[T]) IsOk() bool    { return r.err == nil }
func (r Result[T]) IsError() bool { return r.err != nil }

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(fmt.Sprintf("Result panic: %v", r.err))
	}
	return r.value
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.err != nil {
		return defaultValue
	}
	return r.value
}

func Bind[T any, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	if r.IsError() {
		return Error[U](r.err)
	}
	return fn(r.value)
}

func Map[T any, U any](r Result[T], fn func(T) U) Result[U] {
	if r.IsError() {
		return Error[U](r.err)
	}
	return Ok(fn(r.value))
}

func Tee[T any](r Result[T], fn func(T)) Result[T] {
	if r.IsError() {
		return r
	}
	fn(r.value)
	return r
}

func Try[T any](fn func() (T, error)) Result[T] {
	val, err := fn()
	if err != nil {
		return Error[T](err)
	}
	return Ok(val)
}

func Ensure[T any](r Result[T], condition func(T) bool, err error) Result[T] {
	if r.IsError() {
		return r
	}
	if !condition(r.value) {
		return Error[T](err)
	}
	return r
}

func Match[T any, U any](
	r Result[T],
	onOk func(T) U,
	onErr func(error) U,
) U {

	if r.IsError() {
		return onErr(r.err)
	}

	return onOk(r.value)
}
