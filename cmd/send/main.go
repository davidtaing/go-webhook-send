package main

import (
	"fmt"
	"log"

	server "github.com/davidtaing/go-webhook-send/internal/send-server"
)

func main() {
	s := server.NewServer()

	log.Println("Hello server started on :8080")

	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
