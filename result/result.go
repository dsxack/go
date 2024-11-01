package result

// Result is a generic type that represents the result of an operation that can fail.
// It is either Ok(value) or Error(err).
type Result[T any] struct {
	value T
	err   error
}

// Error creates a new Result with the given error.
func Error[T any](err error) Result[T] { return Result[T]{err: err} }

// Ok creates a new Result with the given value.
func Ok[T any](value T) Result[T] { return Result[T]{value: value} }

// IsErr returns true if the Result is an error.
func (r Result[T]) IsErr() bool { return r.err != nil }

// IsOk returns true if the Result is ok.
func (r Result[T]) IsOk() bool { return r.err == nil }

// Expect panics with the given message if the Result is an error.
// Otherwise, it returns the value.
func (r Result[T]) Expect(msg string) T {
	if r.err != nil {
		panic(msg + ": " + r.err.Error())
	}
	return r.value
}

// ExpectErr panics with the given message if the Result is ok.
// Otherwise, it returns the error.
func (r Result[T]) ExpectErr(msg string) error {
	if r.err == nil {
		panic(msg + ": expected error")
	}
	return r.err
}

// UnwrapErr returns the error of the Result if it is an error.
// It panics if the Result is ok.
func (r Result[T]) UnwrapErr() error {
	if r.err == nil {
		panic("unwrap error called on ok result")
	}
	return r.err
}

// Unwrap returns the value of the Result if it is ok.
// It panics if the Result is an error.
func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic("unwrap called on error: " + r.err.Error())
	}
	return r.value
}

// UnwrapOr returns the value of the Result if it is ok.
// Otherwise, it returns the given default value.
func (r Result[T]) UnwrapOr(def T) T {
	if r.err != nil {
		return def
	}
	return r.value
}

// UnwrapOrDefault returns the value of the Result if it is ok.
// Otherwise, it returns the zero value of the type.
func (r Result[T]) UnwrapOrDefault() T {
	if r.err != nil {
		var def T
		return def
	}
	return r.value
}

// UnwrapOrElse returns the value of the Result if it is ok.
// Otherwise, it returns the result of the given function.
func (r Result[T]) UnwrapOrElse(fn func(error) T) T {
	if r.err != nil {
		return fn(r.err)
	}
	return r.value
}
