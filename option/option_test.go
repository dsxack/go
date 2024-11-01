package option

import "testing"

func TestSome(t *testing.T) {
	expected := 1
	opt := Some(1)

	got := opt.Unwrap()

	if got != expected {
		t.Errorf("Unwrap() = %v; want %v", got, expected)
	}
}

func TestNone(t *testing.T) {
	opt := None[int]()

	if opt.IsSome() {
		t.Errorf("IsSome() = true; want false")
	}

	if !opt.IsNone() {
		t.Errorf("IsNone() = false; want true")
	}

	defaultValue := Option[int]{}
	if opt != defaultValue {
		t.Errorf("None() = %v; want %v", opt, defaultValue)
	}
}

func TestExpect(t *testing.T) {
	expected := 1
	opt := Some(1)

	got := opt.Expect("panic message")

	if got != expected {
		t.Errorf("Expect() = %v; want %v", got, expected)
	}
}

func TestExpectNone(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expect() did not panic")
		}
	}()

	opt := None[int]()
	opt.Expect("panic message")
}

func TestUnwrap(t *testing.T) {
	expected := 1
	opt := Some(1)

	got := opt.Unwrap()

	if got != expected {
		t.Errorf("Unwrap() = %v; want %v", got, expected)
	}
}

func TestUnwrapNone(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Unwrap() did not panic")
		}
	}()

	opt := None[int]()
	opt.Unwrap()
}

func TestUnwrapOr(t *testing.T) {
	expected := 1
	opt := Some(1)

	got := opt.UnwrapOr(2)

	if got != expected {
		t.Errorf("UnwrapOr() = %v; want %v", got, expected)
	}
}

func TestUnwrapOrNone(t *testing.T) {
	expected := 2
	opt := None[int]()

	got := opt.UnwrapOr(2)

	if got != expected {
		t.Errorf("UnwrapOr() = %v; want %v", got, expected)
	}
}

func TestUnwrapOrDefault(t *testing.T) {
	expected := 1
	opt := Some(1)

	got := opt.UnwrapOrDefault()

	if got != expected {
		t.Errorf("UnwrapOrDefault() = %v; want %v", got, expected)
	}
}

func TestUnwrapOrDefaultNone(t *testing.T) {
	var expected int
	opt := None[int]()

	got := opt.UnwrapOrDefault()

	if got != expected {
		t.Errorf("UnwrapOrDefault() = %v; want %v", got, expected)
	}
}

func TestUnwrapOrElse(t *testing.T) {
	expected := 1
	opt := Some(1)

	got := opt.UnwrapOrElse(func() int { return 2 })

	if got != expected {
		t.Errorf("UnwrapOrElse() = %v; want %v", got, expected)
	}
}

func TestUnwrapOrElseNone(t *testing.T) {
	expected := 2
	opt := None[int]()

	got := opt.UnwrapOrElse(func() int { return 2 })

	if got != expected {
		t.Errorf("UnwrapOrElse() = %v; want %v", got, expected)
	}
}
