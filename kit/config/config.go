package config

import (
	"github.com/dsxack/go/v2/config"
	"github.com/google/wire"
)

type ParsedConfig interface {
	Interface
}

type Interface interface {
	config() Config
}

type HTTP struct {
	ListenAddr string
}

type Config struct {
	HTTP
}

func (c Config) config() Config { return c }

func ParseConfig(layer config.Layer, cfg Interface) (ParsedConfig, error) {
	err := config.Parse(layer, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func ProvideConfig(cfg ParsedConfig) Config {
	return cfg.config()
}

var Set = wire.NewSet(
	ParseConfig,
	ProvideConfig,
	wire.FieldsOf(new(Config), "HTTP"),
)
