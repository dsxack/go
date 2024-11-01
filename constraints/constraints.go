package constraints

// This file is a partial copy of the constraints package from the Go experimental repository.
// See https://pkg.go.dev/golang.org/x/exp/constraints for the original implementation.

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}
