package safe

import (
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

// Go starts a recoverable goroutine
func Go(goroutine func()) {
	GoWithRecover(goroutine, defaultRecoverGoroutine)
}

// GoWithRecover starts a recoverable goroutine using given customRecover() function
func GoWithRecover(goroutine func(), customRecover func(err interface{})) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				customRecover(err)
			}
		}()
		goroutine()
	}()
}

func defaultRecoverGoroutine(err interface{}) {
	logrus.Errorf("Error in Go routine: %s", err)
	logrus.Errorf("Stack: %s", debug.Stack())
}
