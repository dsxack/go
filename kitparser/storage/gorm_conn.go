package storage

import (
	"database/sql"
	"github.com/dsxack/go/v2/kitparser/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GORMConnHolder struct {
	Dialector gorm.Dialector
}

func (c *GORMConnHolder) Conn() Conn {
	return c.Dialector
}

type GORMConnFactory struct {
	driverName string
	dsn        string
}

func NewGORMConnFactory(driverName string, dsn string) *GORMConnFactory {
	return &GORMConnFactory{driverName: driverName, dsn: dsn}
}

func NewEnvGORMConnFactory(cfg config.Database) *GORMConnFactory {
	dialect := cfg.Dialect
	if dialect == "" {
		dialect = "sqlite3"
	}
	dsn := cfg.DSN
	if dsn == "" {
		dsn = "database.sqlite"
	}

	return NewGORMConnFactory(dialect, dsn)
}

func (c GORMConnFactory) NewConn() (ConnHolder, func(), error) {
	db, err := sql.Open(c.driverName, c.dsn)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		_ = db.Close()
	}

	var dialector gorm.Dialector
	switch c.driverName {
	case "sqlite3":
		dialector = &sqlite.Dialector{Conn: db}
	case "mysql":
		dialector = mysql.New(mysql.Config{Conn: db})
	default:
		dialector = postgres.New(postgres.Config{Conn: db})
	}

	return &GORMConnHolder{Dialector: dialector}, cleanup, nil
}
