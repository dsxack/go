package http

import (
	kithttp "github.com/dsxack/go/v2/kit/transport/http"
	"github.com/gorilla/mux"
)

type Handler kithttp.Handler

func NewHandler(handler Handler) kithttp.Handler {
	router := mux.NewRouter()
	// TODO: bind parser handlers here.
	router.PathPrefix("/").Handler(handler)
	return kithttp.Handler(router)
}
