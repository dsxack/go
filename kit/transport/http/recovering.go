package http

import (
	"fmt"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

// RecoveringMiddleware recovers panic and returns error to client
func RecoveringMiddleware(handler http.Handler, errorEncoder kithttp.ErrorEncoder) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = fmt.Errorf(t)
				case error:
					err = t
				}

				errorEncoder(request.Context(), err, writer)
			}
		}()

		handler.ServeHTTP(writer, request)
	})
}
