package logger

import (
	"github.com/dsxack/go/v2/kitparser/config"
	"github.com/sirupsen/logrus"
	"os"
)

func NewEnvLogrusLogger(cfg config.Logger) (logrus.FieldLogger, error) {
	levelStr := cfg.Level
	if levelStr == "" {
		levelStr = "debug"
	}

	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		return nil, err
	}

	return &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.TextFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        level,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}, nil
}
