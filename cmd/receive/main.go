package main

import (
	"fmt"
	"log"

	server "github.com/davidtaing/go-webhook-send/internal/receive-server"
)

func main() {
	s := server.NewReceiveServer()

	log.Println("Hello receive server started on :8080")

	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
