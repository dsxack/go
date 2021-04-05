package app

import (
	"github.com/dsxack/go/v2/kit/app"
	kitparserconfig "github.com/dsxack/go/v2/kitparser/config"
	"github.com/dsxack/go/v2/kitparser/logger"
	"github.com/dsxack/go/v2/kitparser/resultlistener"
	"github.com/dsxack/go/v2/kitparser/scraper/http"
	"github.com/dsxack/go/v2/kitparser/session"
	"github.com/dsxack/go/v2/kitparser/storage"
	transportcli "github.com/dsxack/go/v2/kitparser/transport/cli"
	transporthttp "github.com/dsxack/go/v2/kitparser/transport/http"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	app.Set,
	kitparserconfig.Set,
	transporthttp.NewHandler,
	transportcli.NewCommands,
	NewConn,
	NewEventsStorage,
	NewSessionsStorage,
	NewEventsStorageResultsListenerDriver,
	session.Set,
	wire.Bind(new(session.Storage), new(*SessionsStorage)),
)

var EnvSet = wire.NewSet(
	Set,
	logger.NewEnvLogrusLogger,
	http.NewEnvScraper,
	resultlistener.NewEnvListener,
	storage.NewEnvFactory,
	storage.NewEnvConnFactory,
	NewEnvParserStorageFactory,
)

func NewConn(factory storage.ConnFactory) (storage.ConnHolder, func(), error) {
	return factory.NewConn()
}

type StorageFactory struct {
	*storage.Factory
}

func NewEnvParserStorageFactory(conn storage.ConnHolder, cfg kitparserconfig.Database, options ...storage.Option) (*StorageFactory, error) {
	factory, err := storage.NewEnvFactory(conn, nil, cfg, options...)
	if err != nil {
		return nil, err
	}
	return &StorageFactory{Factory: factory}, nil
}

type EventsStorage struct {
	*storage.Storage
}

func NewEventsStorage(factory *StorageFactory) (*EventsStorage, error) {
	eventStorage, err := factory.NewStorage("events", &resultlistener.StorageEvent{})
	if err != nil {
		return nil, err
	}
	return &EventsStorage{Storage: eventStorage}, nil
}

type SessionsStorage struct {
	*storage.Storage
}

func NewSessionsStorage(factory *StorageFactory) (*SessionsStorage, error) {
	sessionsStorage, err := factory.NewStorage("sessions", &session.StorageModel{})
	if err != nil {
		return nil, err
	}
	return &SessionsStorage{Storage: sessionsStorage}, nil
}

func NewEventsStorageResultsListenerDriver(eventsStorage *EventsStorage) *resultlistener.StorageDriver {
	return resultlistener.NewStorageDriver(eventsStorage)
}
