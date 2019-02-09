package http

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	kithttp "github.com/go-kit/kit/transport/http"
)

func TestRecoveringMiddleware(t *testing.T) {
	type args struct {
		handler      http.Handler
		errorEncoder kithttp.ErrorEncoder
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantBody   string
	}{
		{
			name: "test recover panic with string",
			args: args{
				handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					panic("test error message")
				}),
				errorEncoder: kithttp.DefaultErrorEncoder,
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   "test error message",
		},
		{
			name: "test recover panic with error",
			args: args{
				handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					panic(errors.New("test error message"))
				}),
				errorEncoder: kithttp.DefaultErrorEncoder,
			},
			wantStatus: http.StatusInternalServerError,
			wantBody:   "test error message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler := tt.args.handler
			handler = RecoveringMiddleware(handler, tt.args.errorEncoder)

			func() {
				defer func() {
					r := recover()
					assert.Nil(t, r)
					assert.Equal(t, tt.wantStatus, recorder.Code)
					assert.Equal(t, tt.wantBody, recorder.Body.String())
				}()
				handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte{})))
			}()
		})
	}
}
