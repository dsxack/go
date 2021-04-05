package resultlistener

import (
	"context"
	"github.com/dsxack/go/v2/kitparser/session"
	"github.com/google/uuid"
	"log"
)

type Storage interface {
	Put(ctx context.Context, value interface{}) error
}

type StorageEvent struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Session   session.ID
	EventType string
	Level     string
	Err       string
	Values    Values
}

type StorageDriver struct {
	Storage
}

func NewStorageDriver(storage Storage) *StorageDriver {
	return &StorageDriver{Storage: storage}
}

func (s StorageDriver) Info(ctx context.Context, eventType string, values Values) {
	err := s.Storage.Put(ctx, StorageEvent{
		ID:        uuid.New(),
		EventType: eventType,
		Session:   session.From(ctx),
		Level:     "info",
		Values:    values,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func (s StorageDriver) Debug(ctx context.Context, eventType string, values Values) {
	err := s.Storage.Put(ctx, StorageEvent{
		ID:        uuid.New(),
		Session:   session.From(ctx),
		EventType: eventType,
		Level:     "debug",
		Values:    values,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func (s StorageDriver) Warn(ctx context.Context, eventType string, err error, values Values) {
	err = s.Storage.Put(ctx, StorageEvent{
		ID:        uuid.New(),
		Session:   session.From(ctx),
		EventType: eventType,
		Level:     "warn",
		Err:       err.Error(),
		Values:    values,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func (s StorageDriver) Error(ctx context.Context, eventType string, err error, values Values) {
	err = s.Storage.Put(ctx, StorageEvent{
		ID:        uuid.New(),
		Session:   session.From(ctx),
		EventType: eventType,
		Level:     "error",
		Err:       err.Error(),
		Values:    values,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
