package storage

import (
	"context"
	"github.com/dsxack/go/v2/kitparser/config"
	"github.com/dsxack/go/v2/kitparser/resultlistener"
	"github.com/dsxack/go/v2/kitparser/session"
	"github.com/reactivex/rxgo/v2"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type model interface {
	setSession(session.ID)
}

type Model struct {
	Session session.ID `gorm:"column:session"` // TODO: what do if not gorm?
}

func (m *Model) setSession(session session.ID) { m.Session = session }

type Driver interface {
	Migrate(bucketName string, value interface{}) error
	Put(ctx context.Context, bucketName string, value interface{}) error
	Get(ctx context.Context, bucketName string, value interface{}, where ...interface{}) error
	GetIterable(ctx context.Context, bucketName string, value interface{}, where ...interface{}) rxgo.Observable
}

type Option func(s *Storage)

type Storage struct {
	bucketName string
	driver     Driver
	listener   *resultlistener.Listener
}

func NewStorage(bucketName string, driver Driver, options ...Option) *Storage {
	s := &Storage{
		bucketName: bucketName,
		driver:     driver,
	}
	for _, option := range options {
		option(s)
	}
	if s.listener == nil {
		s.listener = resultlistener.NewListener(&resultlistener.NoopDriver{})
	}
	return s
}

func (s Storage) Migrate(value interface{}) error {
	return s.driver.Migrate(s.bucketName, value)
}

const putEventType = "storage:put"
const errEventType = "storage:error"

func (s Storage) Put(ctx context.Context, value interface{}) error {
	err := s.driver.Put(ctx, s.bucketName, value)
	if err != nil {
		s.listener.Error(ctx, errEventType, err, resultlistener.Values{
			"bucketName": s.bucketName,
		})
		return err
	}
	s.listener.Info(ctx, putEventType, resultlistener.Values{
		"bucketName": s.bucketName,
	})
	return nil
}

func (s Storage) Get(ctx context.Context, value interface{}, where ...interface{}) error {
	err := s.driver.Get(ctx, s.bucketName, value, where...)
	if err != nil {
		// TODO: handle error
		return err
	}
	return nil
}

func (s Storage) GetIterable(ctx context.Context, value interface{}, where ...interface{}) rxgo.Observable {
	return s.driver.GetIterable(ctx, s.bucketName, value, where...)
}

func (s Storage) PutMapFunc(ctx context.Context, item interface{}) (interface{}, error) {
	return item, s.Put(ctx, item)
}

func WithListener(listener *resultlistener.Listener) Option {
	return func(s *Storage) {
		s.listener = listener
	}
}

func WithDefaultListener(listener *resultlistener.Listener) Option {
	return func(s *Storage) {
		if s.listener == nil {
			s.listener = listener
		}
	}
}

type Factory struct {
	driver  Driver
	options []Option
}

func NewFactory(driver Driver, options ...Option) *Factory {
	return &Factory{driver: driver, options: options}
}

func NewEnvFactory(
	conn ConnHolder,
	listener *resultlistener.Listener,
	cfg config.Database,
	options ...Option,
) (*Factory, error) {
	// TODO: detect from env.
	gormDialector := conn.Conn().(gorm.Dialector)

	logger := gormlogger.Discard
	if cfg.Debug {
		logger = gormlogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gormlogger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      gormlogger.Info,
			Colorful:      true,
		})
	}

	driver, err := NewSQLDriver(gormDialector, &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	options = append(options, WithDefaultListener(listener))

	return NewFactory(driver, options...), nil
}

func (f Factory) NewStorage(bucketName string, value interface{}) (*Storage, error) {
	storage := NewStorage(bucketName, f.driver, f.options...)
	err := storage.Migrate(value)
	if err != nil {
		return nil, err
	}
	return storage, nil
}
