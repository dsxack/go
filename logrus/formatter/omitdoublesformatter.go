package formatter

import (
	"github.com/sirupsen/logrus"
)

type skipDoublesFormatter struct {
	next        logrus.Formatter
	lastMessage string
}

// NewOmitDoublesFormatter return formatter skips same messages coming in a row
func NewOmitDoublesFormatter(next logrus.Formatter) *skipDoublesFormatter {
	return &skipDoublesFormatter{next: next}
}

func (f *skipDoublesFormatter) Format(e *logrus.Entry) ([]byte, error) {
	if f.lastMessage == e.Message {
		return []byte{}, nil
	}
	f.lastMessage = e.Message
	return f.next.Format(e)
}
