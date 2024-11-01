package option

// Option is a type that represents an optional value.
// It is either Some(value) or None.
// The Option type is generic and can hold any type of value.
// The zero value of Option is None.
type Option[T any] struct {
	value T
	isSet bool
}

// Some creates a new Option with the given value.
func Some[T any](value T) Option[T] { return Option[T]{value: value, isSet: true} }

// None creates a new Option with no value.
func None[T any]() Option[T] { return Option[T]{} }

// IsSome returns true if the Option is Some.
func (o Option[T]) IsSome() bool { return o.isSet }

// IsNone returns true if the Option is None.
func (o Option[T]) IsNone() bool { return !o.isSet }

// Expect panics with the given message if the Option is None.
// Otherwise, it returns the value.
func (o Option[T]) Expect(msg string) T {
	if !o.isSet {
		panic(msg)
	}
	return o.value
}

// Unwrap returns the value of the Option if it is Some.
// It panics if the Option is None.
func (o Option[T]) Unwrap() T {
	if !o.isSet {
		panic("unwrap called on None")
	}
	return o.value
}

// UnwrapOr returns the value of the Option if it is Some.
// Otherwise, it returns the given default value.
func (o Option[T]) UnwrapOr(def T) T {
	if !o.isSet {
		return def
	}
	return o.value
}

// UnwrapOrDefault returns the value of the Option if it is Some.
// Otherwise, it returns the zero value of the type.
func (o Option[T]) UnwrapOrDefault() T {
	if !o.isSet {
		var def T
		return def
	}
	return o.value
}

// UnwrapOrElse returns the value of the Option if it is Some.
// Otherwise, it returns the result of the given function.
func (o Option[T]) UnwrapOrElse(fn func() T) T {
	if !o.isSet {
		return fn()
	}
	return o.value
}
