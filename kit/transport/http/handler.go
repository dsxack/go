package http

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

type Handler http.Handler

type kitEndpoint = func(ctx context.Context, request interface{}) (response interface{}, err error)

func NewServer(
	e kitEndpoint,
	dec kithttp.DecodeRequestFunc,
	enc kithttp.EncodeResponseFunc,
	options ...kithttp.ServerOption,
) *kithttp.Server {
	return kithttp.NewServer(e, dec, enc, options...)
}
