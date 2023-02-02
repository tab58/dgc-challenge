package main

import (
	"github.com/tab58/dgc-challenge/internal/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		panic(err)
	}
	srv.Start()
}
