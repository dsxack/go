package safe

import "testing"

func TestGo(t *testing.T) {
	Go(func() {
		panic("error")
	})
}
