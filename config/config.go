package config

import "github.com/mitchellh/mapstructure"

func Parse(layer Layer, cfg interface{}) error {
	values, err := layer.Values()
	if err != nil {
		return err
	}

	return mapstructure.Decode(values, cfg)
}
