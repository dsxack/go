package session

import (
	"context"
	"database/sql/driver"
	"github.com/google/wire"
)

type Storage interface {
	Put(ctx context.Context, value interface{}) error
	Get(ctx context.Context, value interface{}, where ...interface{}) error
}

type Status string

func (s Status) Value() (driver.Value, error) {
	return []byte(s), nil
}

func (s *Status) Scan(src interface{}) error {
	*s = Status(src.([]byte))
	return nil
}

const (
	CompletedStatus Status = "completed"
	NewStatus       Status = "new"
)

func (s Status) IsCompleted() bool {
	return s == CompletedStatus
}

func (s Status) Complete() Status {
	return CompletedStatus
}

func (s Status) New() Status {
	return NewStatus
}

// TODO: put into domain?[
type StorageModel struct {
	ID     ID `gorm:"primaryKey"`
	Status Status
}

type Repository struct {
	Storage Storage
}

func (s Repository) New(ctx context.Context) (*ID, context.Context, error) {
	id := newID()
	err := s.Storage.Put(ctx, StorageModel{
		ID:     id,
		Status: NewStatus,
	})
	if err != nil {
		return nil, nil, err
	}
	return &id, With(ctx, id), nil
}

func (s Repository) Complete(ctx context.Context, id ID) error {
	return s.Storage.Put(ctx, StorageModel{
		ID:     id,
		Status: CompletedStatus,
	})
}

func (s Repository) Get(ctx context.Context, id ID) (*StorageModel, error) {
	var model StorageModel
	err := s.Storage.Get(ctx, &model, "id=?", id)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

var Set = wire.NewSet(
	wire.Struct(new(Repository), "*"),
)
