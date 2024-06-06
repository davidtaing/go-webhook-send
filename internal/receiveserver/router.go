package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/davidtaing/go-webhook-send/internal/server"
)

type RouteHandler struct {
}

func (rh *RouteHandler) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", server.MakeHTTPHandler(rh.HelloWorldHandler))
	r.HandleFunc("/receive", server.MakeHTTPHandler(rh.Receive))

	return r
}

// re-export server.MakeHTTPHandler to make testing easier
func (rh *RouteHandler) MakeHTTPHandler(f server.ApiFunc) http.HandlerFunc {
	return server.MakeHTTPHandler(f)
}
