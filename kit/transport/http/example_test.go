package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

type (
	someRequest struct{}
)

func makeSomeEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, nil
	}
}

var (
	request        *http.Request
	responseWriter http.ResponseWriter
)

func ExampleMakeJSONRequestDecoder() {
	kithttp.NewServer(
		makeSomeEndpoint(),
		NewJSONRequestDecoder(func() interface{} { return &someRequest{} }),
		kithttp.EncodeJSONResponse,
	)
}

func ExampleRecoveringMiddleware() {
	var handler http.Handler
	{
		handler = kithttp.NewServer(
			makeSomeEndpoint(),
			NewJSONRequestDecoder(func() interface{} { return &someRequest{} }),
			kithttp.EncodeJSONResponse,
		)
		handler = RecoveringMiddleware(handler, kithttp.DefaultErrorEncoder)
	}

	// do something with handler
	handler.ServeHTTP(responseWriter, request)
}
