package logger

import (
	"errors"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/sirupsen/logrus"
)

type LogrusAdapterOption func(*LogrusAdapter)

func MessageKey(key string) LogrusAdapterOption {
	return func(a *LogrusAdapter) { a.messageKey = key }
}

func ErrorKey(key string) LogrusAdapterOption {
	return func(a *LogrusAdapter) { a.errorKey = key }
}

type LogrusAdapter struct {
	logrus.FieldLogger
	messageKey string
	errorKey   string
}

var errMissingValue = errors.New("(MISSING)")

func NewLogrusAdapter(logger logrus.FieldLogger, opts ...LogrusAdapterOption) *LogrusAdapter {
	adapter := &LogrusAdapter{
		FieldLogger: logger,
	}
	for _, opt := range opts {
		opt(adapter)
	}
	return adapter
}

func (l LogrusAdapter) Log(keyvals ...interface{}) error {
	fields := logrus.Fields{}

	var errorValue error
	var messageValue interface{}
	var levelValue level.Value

	for i := 0; i < len(keyvals); i += 2 {
		var key = fmt.Sprint(keyvals[i])
		var value interface{}

		if i+1 < len(keyvals) {
			value = keyvals[i+1]
			fields[key] = value
		} else {
			fields[key] = errMissingValue
		}

		switch key {
		case l.messageKey:
			messageValue = value
			delete(fields, key)
		case l.errorKey:
			if errVal, ok := value.(error); ok {
				errorValue = errVal
			}
			delete(fields, key)
		case level.Key():
			delete(fields, key)
		}
	}

	logger := l.WithFields(fields)
	if errorValue != nil {
		logger = logger.WithError(errorValue)
	}

	var loggerFunc func(...interface{})
	switch levelValue {
	case level.DebugValue():
		loggerFunc = logger.Debug
	case level.InfoValue():
		loggerFunc = logger.Info
	case level.WarnValue():
		loggerFunc = logger.Warn
	case level.ErrorValue():
		loggerFunc = logger.Error
	default:
		loggerFunc = logger.Debug
	}

	if messageValue != nil {
		loggerFunc(messageValue)
	} else {
		loggerFunc()
	}

	return nil
}
