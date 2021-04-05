package logger

import (
	"context"
	"github.com/dsxack/go/v2/kitparser/session"
	"github.com/sirupsen/logrus"
)

const SessionKey = "session"

func PopulateByContext(logger logrus.FieldLogger, ctx context.Context) logrus.FieldLogger {
	return logger.WithFields(logrus.Fields{
		SessionKey: session.From(ctx),
	})
}
