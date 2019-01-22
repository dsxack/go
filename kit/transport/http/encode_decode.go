package http

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
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
