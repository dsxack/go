package resultlistener

import (
	"context"
	kitparserlogger "github.com/dsxack/go/v2/kitparser/logger"
	"github.com/dsxack/go/v2/kitparser/session"
	"github.com/sirupsen/logrus"
)

const ValuesKey = "values"

type LogDriver struct {
	logger logrus.FieldLogger
}

func (l LogDriver) Info(ctx context.Context, eventType string, values Values) {
	l.logger.WithFields(logrus.Fields{
		kitparserlogger.SessionKey: session.From(ctx),
		ValuesKey:                  values,
	}).Info(eventType)
}

func (l LogDriver) Debug(ctx context.Context, eventType string, values Values) {
	l.logger.WithFields(logrus.Fields{
		kitparserlogger.SessionKey: session.From(ctx),
		ValuesKey:                  values,
	}).Debug(eventType)
}

func (l LogDriver) Warn(ctx context.Context, eventType string, err error, values Values) {
	l.logger.WithFields(logrus.Fields{
		kitparserlogger.SessionKey: session.From(ctx),
		ValuesKey:                  values,
	}).WithError(err).Warn(eventType)
}

func (l LogDriver) Error(ctx context.Context, eventType string, err error, values Values) {
	l.logger.WithFields(logrus.Fields{
		kitparserlogger.SessionKey: session.From(ctx),
		ValuesKey:                  values,
	}).WithError(err).Error(eventType)
}

func NewLogDriver(logger logrus.FieldLogger) *LogDriver {
	return &LogDriver{logger: logger}
}
