package http

import (
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

type (
	someRequest struct{}
)

var (
	makeSomeEndpoint = func() endpoint.Endpoint {}
)

func ExampleMakeJSONRequestDecoder() {
	kithttp.NewServer(
		makeSomeEndpoint(),
		MakeJSONRequestDecoder(func() interface{} { return &someRequest{} }),
		kithttp.EncodeJSONResponse,
	)
}

func ExampleRecoveringMiddleware() {
	var handler http.Handler
	{
		handler = kithttp.NewServer(
			makeSomeEndpoint(),
			MakeJSONRequestDecoder(func() interface{} { return &someRequest{} }),
			kithttp.EncodeJSONResponse,
		)
		handler = RecoveringMiddleware(handler, kithttp.DefaultErrorEncoder)
	}
}
