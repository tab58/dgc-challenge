package server

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/tab58/dgc-challenge/internal/api/db"
)

type server struct {
	srv  *echo.Echo
	port int

	dbClient db.DBClient

	controller Controller
}

func (s *server) Start() {
	listenPort := fmt.Sprintf(":%d", s.port)
	s.srv.Logger.Fatal(s.srv.Start(listenPort))
}

func NewServer(config *Config) (*server, error) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	dbClient, err := db.NewSomeDBClient()
	if err != nil {
		return nil, err
	}

	controller, err := NewController(dbClient)
	if err != nil {
		return nil, err
	}

	server := &server{
		srv:  e,
		port: config.Port,

		controller: controller,
	}
	AttachRoutes(server)
	return server, nil
}
