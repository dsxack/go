package resultlistener

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type Event struct {
	Type   string            `json:"type"`
	Values map[string]string `json:"values"`
	Err    error             `json:"err,omitempty"`
}

type Values map[string]string

func (v Values) Value() (driver.Value, error) {
	return json.Marshal(v)
}

func (v *Values) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), v)
}

func (e Event) Failed() error { return e.Err }

type Driver interface {
	Info(ctx context.Context, eventType string, values Values)
	Warn(ctx context.Context, eventType string, err error, values Values)
	Error(ctx context.Context, eventType string, err error, values Values)
	Debug(ctx context.Context, eventType string, values Values)
}

type Listener struct {
	driver Driver
}

func NewListener(driver Driver) *Listener {
	return &Listener{driver: driver}
}

func (l Listener) Info(ctx context.Context, eventType string, values Values) {
	l.driver.Info(ctx, eventType, values)
}

func (l Listener) Error(ctx context.Context, eventType string, err error, values Values) {
	l.driver.Error(ctx, eventType, err, values)
}

func (l Listener) Warn(ctx context.Context, eventType string, err error, values Values) {
	l.driver.Warn(ctx, eventType, err, values)
}

func (l Listener) Debug(ctx context.Context, eventType string, values Values) {
	l.driver.Debug(ctx, eventType, values)
}

func NewEnvListener(logger logrus.FieldLogger, storageDriver *StorageDriver) *Listener {
	return NewListener(NewComposeDriver(
		NewLogDriver(logger),
		storageDriver,
	))
}
