package config

import (
	kitconfig "github.com/dsxack/go/v2/kit/config"
	"github.com/google/wire"
	"time"
)

// TODO: add available custom parserConfig from concrete parser

type Interface interface {
	config() Config
}

type Database struct {
	Dialect string
	DSN     string
	Debug   bool
}

type Logger struct {
	Level string
}

type ScraperHTTP struct {
	Parallelism int
	RandomDelay time.Duration
}

type Config struct {
	kitconfig.Config `mapstructure:",squash"`
	Database
	Logger
	ScraperHTTP
}

func (c Config) config() Config { return c }

func ProvideConfig(cfg kitconfig.ParsedConfig) Config {
	//goland:noinspection GoVetImpossibleInterfaceToInterfaceAssertion
	return cfg.(Interface).config()
}

var Set = wire.NewSet(
	ProvideConfig,
	wire.FieldsOf(new(Config), "Database", "ScraperHTTP", "Logger"),
)
