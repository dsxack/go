package resultlistener

import "context"

type NoopDriver struct{}

func NewNoopDriver() *NoopDriver {
	return &NoopDriver{}
}
func (d NoopDriver) Info(_ context.Context, _ string, _ Values)           {}
func (d NoopDriver) Warn(_ context.Context, _ string, _ error, _ Values)  {}
func (d NoopDriver) Error(_ context.Context, _ string, _ error, _ Values) {}
func (d NoopDriver) Debug(_ context.Context, _ string, _ Values)          {}
