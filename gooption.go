package gooption

type Option[T any] struct {
	Value *T
}

func (o Option[T]) IsSome() bool {
	return o.Value != nil
}

func (o Option[T]) IsNone() bool {
	return o.Value == nil
}

func (o Option[T]) Unwrap() T {
	if o.IsSome() {
		return *o.Value
	}
	panic("called `Unwrap()` on a `None` value")
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.IsSome() {
		return *o.Value
	}
	return def
}

func (o Option[T]) UnwrapOrDefault() T {
	if o.IsSome() {
		return *o.Value
	}
	var empty T
	return empty
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.IsSome() {
		return *o.Value
	}
	return f()
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		Value: &value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		Value: nil,
	}
}
