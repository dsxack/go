package result

import (
	"errors"
	"fmt"
	"testing"
)

func TestOk(t *testing.T) {
	result := Ok(1)
	if result.IsErr() {
		t.Errorf("Ok(1).IsErr() = true; want false")
	}
	if !result.IsOk() {
		t.Errorf("Ok(1).IsOk() = false; want true")
	}
}

func TestError(t *testing.T) {
	err := fmt.Errorf("some error")
	result := Error[int](err)
	if !result.IsErr() {
		t.Errorf("Error(err).IsErr() = false; want true")
	}
	if result.IsOk() {
		t.Errorf("Error(err).IsOk() = true; want false")
	}
}

func TestExpect(t *testing.T) {
	result := Ok(1)
	got := result.Expect("panic message")
	if got != 1 {
		t.Errorf("Ok(1).Expect() = %v; want 1", got)
	}
}

func TestExpectPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expect() did not panic")
		}
	}()

	result := Error[int](fmt.Errorf("some error"))
	_ = result.Expect("panic message")
}

func TestExpectErr(t *testing.T) {
	err := fmt.Errorf("some error")
	result := Error[int](err)
	got := result.ExpectErr("panic message")
	if !errors.Is(got, err) {
		t.Errorf("Error(err).ExpectErr() = %v; want %v", got, err)
	}
}

func TestExpectErrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ExpectErr() did not panic")
		}
	}()

	result := Ok(1)
	_ = result.ExpectErr("panic message")
}

func TestUnwrapErr(t *testing.T) {
	err := fmt.Errorf("some error")
	result := Error[int](err)

	got := result.UnwrapErr()
	if !errors.Is(got, err) {
		t.Errorf("Error(err).UnwrapErr() = %v; want %v", got, err)
	}
}

func TestUnwrapErrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("UnwrapErr() did not panic")
		}
	}()

	result := Ok(1)
	_ = result.UnwrapErr()
}

func TestUnwrap(t *testing.T) {
	result := Ok(1)
	got := result.Unwrap()
	if got != 1 {
		t.Errorf("Ok(1).Unwrap() = %v; want 1", got)
	}
}

func TestUnwrapPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Unwrap() did not panic")
		}
	}()

	err := fmt.Errorf("some error")
	result := Error[int](err)
	_ = result.Unwrap()
}

func TestUnwrapOr(t *testing.T) {
	result := Ok(1)
	got := result.UnwrapOr(2)
	if got != 1 {
		t.Errorf("Ok(1).UnwrapOr(2) = %v; want 1", got)
	}

}

func TestUnwrapOrError(t *testing.T) {
	result := Error[int](fmt.Errorf("some error"))
	got := result.UnwrapOr(2)
	if got != 2 {
		t.Errorf("Error(err).UnwrapOr(2) = %v; want 2", got)
	}
}

func TestUnwrapOrDefault(t *testing.T) {
	result := Ok(1)
	got := result.UnwrapOrDefault()
	if got != 1 {
		t.Errorf("Ok(1).UnwrapOrDefault() = %v; want 1", got)
	}
}

func TestUnwrapOrDefaultError(t *testing.T) {
	result := Error[int](fmt.Errorf("some error"))
	var expected int
	got := result.UnwrapOrDefault()
	if got != expected {
		t.Errorf("Error(err).UnwrapOrDefault() = %v; want %v", got, expected)
	}
}

func TestUnwrapOrElse(t *testing.T) {
	result := Ok(1)
	got := result.UnwrapOrElse(func(err error) int {
		return 2
	})
	if got != 1 {
		t.Errorf("Ok(1).UnwrapOrElse() = %v; want 1", got)
	}
}

func TestUnwrapOrElseError(t *testing.T) {
	result := Error[int](fmt.Errorf("some error"))
	got := result.UnwrapOrElse(func(err error) int {
		return 2
	})
	if got != 2 {
		t.Errorf("Error(err).UnwrapOrElse() = %v; want 2", got)
	}
}
