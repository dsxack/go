package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type rawResponderResponse struct {
}

func (rawResponderResponse) StatusCode() int {
	return http.StatusAccepted
}

func (rawResponderResponse) RawResponse() io.Reader {
	return strings.NewReader("test raw message")
}

type plainFailerResponse struct {
}

func (plainFailerResponse) Failed() error {
	return errors.New("test plain error message")
}

type jsonFailerResponse struct {
}

type jsonError struct {
	error
	statusCode int
}

func (e jsonError) StatusCode() int {
	return e.statusCode
}

func (e jsonError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Error string `json:"error"`
	}{
		Error: e.Error(),
	})
}

func (jsonFailerResponse) Failed() error {
	return jsonError{
		error:      errors.New("test json error message"),
		statusCode: http.StatusBadRequest,
	}
}

type statusCreatedResponse struct {
}

func (statusCreatedResponse) StatusCode() int {
	return http.StatusCreated
}

type statusNoContentResponse struct {
}

func (statusNoContentResponse) StatusCode() int {
	return http.StatusNoContent
}

type headererResponse struct {
}

func (headererResponse) Headers() http.Header {
	h := http.Header{}
	h.Add("Location", "test location header")
	return h
}

func TestEncodeJSONResponse(t *testing.T) {
	tests := []struct {
		name       string
		response   interface{}
		assertFunc func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "test raw response",
			response: rawResponderResponse{},
			assertFunc: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, "test raw message", recorder.Body.String())
				assert.Equal(t, http.StatusAccepted, recorder.Code)
			},
		},
		{
			name:     "test plain failer response",
			response: plainFailerResponse{},
			assertFunc: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
				assert.Equal(t, "test plain error message", recorder.Body.String())
				assert.Equal(t, recorder.Header().Get("Content-Type"), "text/plain; charset=utf-8")
			},
		},
		{
			name:     "test json failer response",
			response: jsonFailerResponse{},
			assertFunc: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
				assert.Equal(t, `{"error":"test json error message"}`, recorder.Body.String())
				assert.Equal(t, recorder.Header().Get("Content-Type"), "application/json; charset=utf-8")
			},
		},
		{
			name:     "test simple status coder response",
			response: statusCreatedResponse{},
			assertFunc: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, recorder.Code, http.StatusCreated)
				assert.Equal(t, recorder.Body.String(), "{}\n")
			},
		},
		{
			name:     "test no content status coder response",
			response: statusNoContentResponse{},
			assertFunc: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, recorder.Code, http.StatusNoContent)
				assert.Equal(t, recorder.Body.String(), "")
			},
		},
		{
			name:     "test headerer response",
			response: headererResponse{},
			assertFunc: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, recorder.Header().Get("Location"), "test location header")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			err := EncodeJSONResponse(context.Background(), recorder, tt.response)
			if err != nil {
				t.Fatal(err)
			}

			tt.assertFunc(t, recorder)
		})
	}
}

type testRequest struct {
	Message string `json:"message"`
}

func TestMakeJSONRequestDecoder(t *testing.T) {
	decoder := NewJSONRequestDecoder(func() interface{} {
		return &testRequest{}
	})

	req, err := decoder(
		context.Background(),
		httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"message":"test request message"}`)),
	)
	assert.Nil(t, err)

	testReq, ok := req.(testRequest)
	if !ok {
		t.Fatal(errors.New("error assert request"))
	}

	assert.Equal(t, testReq.Message, "test request message")
}

func TestMakeJSONRequestDecoderError(t *testing.T) {
	decoder := NewJSONRequestDecoder(func() interface{} {
		return &testRequest{}
	})

	_, err := decoder(
		context.Background(),
		httptest.NewRequest(http.MethodPost, "/", nil),
	)
	assert.NotNil(t, err)
}
