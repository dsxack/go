package logrus

import (
	"github.com/dsxack/go/v2/safe"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

func init() {
	safe.DefaultRecover = RecoverLogger
}

func RecoverLogger(err interface{}) {
	logrus.Errorf("Error in Go routine: %s\nStack: %s", err, debug.Stack())
}
