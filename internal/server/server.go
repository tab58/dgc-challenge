package server

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type server struct {
	srv  *echo.Echo
	port int

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

	controller, err := NewController()
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
