package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/tab58/dgc-challenge/internal/server"
)

func main() {
	envPtr := flag.String("env", "dev", "the configuration environment for the server")
	env := strings.ToLower(*envPtr)
	switch env {
	case "dev":
		fmt.Println("Running in dev mode...")
	case "prod":
		fmt.Println("Running in production mode...")
	default:
		panic(fmt.Errorf("Invalid environment provided: '" + env + "'"))
	}

	// get the proper config for the server
	cfg, err := server.NewConfig(env)
	if err != nil {
		panic(err)
	}

	// create the HTTP server
	srv, err := server.NewServer(cfg)
	if err != nil {
		panic(err)
	}
	srv.Start()
}
