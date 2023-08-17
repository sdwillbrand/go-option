package goopt

type Option[T any] struct {
	value *T
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func (o Option[T]) Unwrap() T {
	if o.value != nil {
		return *o.value
	} else {
		panic("")
	}
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.value != nil {
		return *o.value
	} else {
		return def
	}
}
func (o Option[T]) UnwrapOrDefault() T {
	if o.value != nil {
		return *o.value
	} else {
		var empty T
		return empty
	}
}

func (o Option[T]) Expect(err string) T {
	if o.value != nil {
		return *o.value
	} else {
		panic(err)
	}
}

func (o *Option[T]) Replace(value T) Option[T] {
	o.value = &value
	return *o
}

func (o Option[T]) Clone() Option[T] {
	clone := *o.value
	return Option[T]{value: &clone}
}

func (o Option[T]) IsSome() bool {
	return o.value != nil
}

func (o Option[T]) IsSomeAnd(predicate func(T) bool) bool {
	return o.value != nil && predicate(*o.value)
}

func (o Option[T]) IsNone() bool {
	return o.value == nil
}
