package session

import (
	"context"
	"github.com/google/uuid"
)

type contextKeyType string

const contextKey = contextKeyType("session")

type ID struct {
	uuid.UUID
}

func newID() ID {
	return ID{
		UUID: uuid.Must(uuid.NewRandom()),
	}
}

func With(ctx context.Context, id ID) context.Context {
	return context.WithValue(ctx, contextKey, id)
}

func From(ctx context.Context) ID {
	return ctx.Value(contextKey).(ID)
}
