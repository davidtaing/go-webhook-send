package server

import (
	"net/http"

	"github.com/davidtaing/go-webhook-send/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func NewSendServer() *http.Server {
	rh := &RouteHandler{}

	return server.NewServer(rh.RegisterRoutes())
}
