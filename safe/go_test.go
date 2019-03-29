package safe

import (
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	Go(func() {
		panic("error")
	})

	time.Sleep(time.Second)
}
