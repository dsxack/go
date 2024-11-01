package safe

import (
	"log"
	"runtime/debug"
)

var DefaultRecover = func(err interface{}) {
	log.Printf("Error in goroutine: %s\nStack: %s\n", err, debug.Stack())
}

// Go starts a recoverable goroutine
func Go(goroutine func()) {
	GoWithRecover(goroutine, DefaultRecover)
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
