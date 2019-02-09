package http

import (
	"context"
	"encoding/json"
	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"io"
	"net/http"
	"reflect"
)

// MakeJSONRequestDecoder is the universal DecodeRequestFunc creator
func MakeJSONRequestDecoder(requestFactory func() interface{}) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		request := requestFactory()
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			return nil, err
		}

		refReq := reflect.ValueOf(request)
		return refReq.Elem().Interface(), nil
	}
}

// RawResponder is checked by EncodeJSONResponse. It should return io.Reader to copy it to http.ResponseWriter
type RawResponder interface {
	RawResponse() io.Reader
}

// EncodeJSONResponse is a EncodeResponseFunc that serializes the response as a
// JSON object to the ResponseWriter. Its works like standard go-lit EncodeJSONResponse but
// checks Failer and RawResponder interfaces
func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if failer, ok := response.(kitendpoint.Failer); ok && failer.Failed() != nil {
		kithttp.DefaultErrorEncoder(ctx, failer.Failed(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if headerer, ok := response.(kithttp.Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}

	code := http.StatusOK
	if sc, ok := response.(kithttp.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	if code == http.StatusNoContent {
		return nil
	}

	if rawResponse, ok := response.(RawResponder); ok {
		_, err := io.Copy(w, rawResponse.RawResponse())
		return err
	}

	return json.NewEncoder(w).Encode(response)
}
