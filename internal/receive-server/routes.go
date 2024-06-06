package server

import (
	"net/http"

	"github.com/davidtaing/go-webhook-send/internal/server"
)

func (rh *RouteHandler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) error {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	return server.WriteJSON(w, http.StatusOK, resp)
}

func (rh *RouteHandler) Receive(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return server.ApiError{Err: http.StatusText(http.StatusMethodNotAllowed), Status: http.StatusMethodNotAllowed}
	}

	resp := make(map[string]string)
	resp["message"] = "Hello World"

	return server.WriteJSON(w, http.StatusOK, resp)
}
