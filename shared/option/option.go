package option

type Option[T any] struct {
	value T
	ok    bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: value,
		ok:    true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsSome() bool {
	return o.ok
}

func (o Option[T]) IsNone() bool {
	return !o.ok
}

func (o Option[T]) Unwrap() T {
	if !o.ok {
		panic("option: unwrap called on none")
	}
	return o.value
}

func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.ok {
		return o.value
	}
	return defaultValue
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	if o.ok {
		return o.value
	}
	return fn()
}

func (o Option[T]) Value() (T, bool) {
	return o.value, o.ok
}

func Map[T any, U any](
	o Option[T],
	fn func(T) U,
) Option[U] {

	if o.IsNone() {
		return None[U]()
	}

	return Some(fn(o.value))
}

func Bind[T any, U any](
	o Option[T],
	fn func(T) Option[U],
) Option[U] {

	if o.IsNone() {
		return None[U]()
	}

	return fn(o.value)
}

func Tee[T any](
	o Option[T],
	fn func(T),
) Option[T] {

	if o.IsSome() {
		fn(o.value)
	}

	return o
}

func Filter[T any](
	o Option[T],
	predicate func(T) bool,
) Option[T] {

	if o.IsNone() {
		return o
	}

	if predicate(o.value) {
		return o
	}

	return None[T]()
}

func Match[T any, U any](
	o Option[T],
	onSome func(T) U,
	onNone func() U,
) U {

	if o.IsSome() {
		return onSome(o.value)
	}

	return onNone()
}
