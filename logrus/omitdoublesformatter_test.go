package logrus_test

import (
	"bytes"
	_logrus "github.com/dsxack/go/logrus"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_skipDoublesFormatter_Format(t *testing.T) {
	formatter := _logrus.NewOmitDoublesFormatter(&logrus.JSONFormatter{})
	logger := logrus.New()
	logger.SetFormatter(formatter)

	var buf bytes.Buffer

	logger.SetOutput(&buf)
	now, _ := time.Parse(time.RFC3339, "2019-03-29T05:13:24+03:00")

	logger.WithTime(now).Error("test")
	logger.WithTime(now).Error("test")

	assert.Equal(t, buf.String(), `{"level":"error","msg":"test","time":"2019-03-29T05:13:24+03:00"}`+"\n")
}
