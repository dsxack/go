package config

import (
	"os"
	"strings"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

type Values map[string]interface{}

type Layer interface {
	Values() (Values, error)
}

type YAMLLayer struct {
	Source Source
}

func NewYAMLLayer(source Source) *YAMLLayer {
	return &YAMLLayer{Source: source}
}

func (l YAMLLayer) Values() (values Values, err error) {
	reader, err := l.Source.Reader()
	if err != nil {
		return values, err
	}

	err = yaml.NewDecoder(reader).Decode(&values)
	return values, err
}

type MergeLayer struct {
	layers    []Layer
	skipError bool
}

func NewMergeLayer(layers ...Layer) *MergeLayer {
	return &MergeLayer{layers: layers}
}

func NewMergeLayerWithSkipError(layers ...Layer) *MergeLayer {
	return &MergeLayer{
		layers:    layers,
		skipError: true,
	}
}

func (m MergeLayer) Values() (Values, error) {
	result := Values{}

	for _, layer := range m.layers {
		layerValues, err := layer.Values()
		if err != nil {
			if m.skipError {
				continue
			}
			return result, err
		}

		err = mergo.Merge(&result, layerValues, mergo.WithOverride)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

type EnvLayer struct {
	prefix string
}

func NewEnvLayer(prefix string) *EnvLayer {
	return &EnvLayer{prefix: prefix}
}

func (e EnvLayer) Values() (Values, error) {
	result := Values{}

	for _, s := range os.Environ() {
		fields := strings.Split(s, "=")
		if len(fields) < 2 {
			continue
		}

		name := fields[0]
		if !strings.HasPrefix(name, e.prefix) {
			continue
		}

		name = strings.TrimPrefix(name, e.prefix)
		name = strings.ToLower(name)

		value := strings.Join(fields[1:], "=")
		recursiveSet(result, strings.Split(name, "_"), value)
	}

	return result, nil
}

func recursiveSet(m map[string]interface{}, path []string, value interface{}) {
	if len(path) == 1 {
		m[path[0]] = value
		return
	}
	sub, ok := m[path[0]]
	if !ok {
		sub = map[string]interface{}{}
		m[path[0]] = sub
	}
	recursiveSet(sub.(map[string]interface{}), path[1:], value)
}
