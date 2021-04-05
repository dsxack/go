package storage

import "github.com/dsxack/go/v2/kitparser/config"

type ConnHolder interface {
	Conn() Conn
}

type Conn interface{}

type ConnFactory interface {
	NewConn() (ConnHolder, func(), error)
}

func NewEnvConnFactory(cfg config.Database) ConnFactory {
	// TODO: detect from env
	return NewEnvGORMConnFactory(cfg)
}
