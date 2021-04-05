package config

import (
	"bytes"
	"io"
	"io/fs"
)

type Source interface {
	Reader() (io.Reader, error)
}

type SourceFunc func() (io.Reader, error)

func (f SourceFunc) Reader() (io.Reader, error) {
	return f()
}

func NewBytesSource(source []byte) Source {
	return SourceFunc(func() (io.Reader, error) {
		return bytes.NewReader(source), nil
	})
}

type FSSource struct {
	FS   fs.FS
	Name string
}

func NewFSSource(fs fs.FS, name string) *FSSource {
	return &FSSource{FS: fs, Name: name}
}

func (s FSSource) Reader() (io.Reader, error) {
	return s.FS.Open(s.Name)
}

type FallbackSource struct {
	Source         Source
	FallbackSource Source
}

func NewFallbackSource(source Source, fallbackSource Source) *FallbackSource {
	return &FallbackSource{Source: source, FallbackSource: fallbackSource}
}

func (s FallbackSource) Reader() (io.Reader, error) {
	r, err := s.Source.Reader()
	if err == nil {
		return r, nil
	}

	return s.FallbackSource.Reader()
}
