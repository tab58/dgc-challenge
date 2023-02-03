package server

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/tab58/dgc-challenge/internal/api"
)

// Attaches the routes to the server.
func AttachRoutes(s *server) {
	s.srv.GET("/books",
		s.HandleGetBookList(),
		ValidateParams(api.GetBookListRequestParams{}),
	)

	s.srv.POST("/books",
		s.HandleCreateBook(),
		ValidateParams(api.GetBookRequestParams{}),
	)
	s.srv.GET("/books/:id",
		s.HandleGetBook(),
		ValidateParams(api.GetBookRequestParams{}),
	)
	s.srv.PUT("/books/:id",
		s.HandleUpdateBook(),
		ValidateParams(api.UpdateBookRequestParams{}),
	)
	s.srv.DELETE("/books/:id",
		s.HandleDeleteBook(),
		ValidateParams(api.DeleteBookRequestParams{}),
	)
}

func (s *server) HandleGetBookList() echo.HandlerFunc {
	handler, err := NewDefaultRouteHandler[api.GetBookListRequestParams, api.GetBookListResponse](
		WithRequestParamType(api.GetBookListRequestParams{}),
	)
	if err != nil {
		panic(err)
	}
	return handler(s.controller.GetBookList)
}

func (s *server) HandleCreateBook() echo.HandlerFunc {
	handler, err := NewDefaultRouteHandler[api.CreateBookRequestParams, api.CreateBookResponse](
		WithRequestParamType(api.GetBookListRequestParams{}),
		WithSuccessCode(http.StatusCreated),
	)
	if err != nil {
		panic(err)
	}
	return handler(s.controller.CreateBook)
}

func (s *server) HandleGetBook() echo.HandlerFunc {
	handler, err := NewDefaultRouteHandler[api.GetBookRequestParams, api.GetBookResponse](
		WithRequestParamType(api.GetBookRequestParams{}),
	)
	if err != nil {
		panic(err)
	}
	return handler(s.controller.GetBook)
}

func (s *server) HandleUpdateBook() echo.HandlerFunc {
	handler, err := NewDefaultRouteHandler[api.UpdateBookRequestParams, api.UpdateBookResponse](
		WithRequestParamType(api.UpdateBookRequestParams{}),
	)
	if err != nil {
		panic(err)
	}
	return handler(s.controller.UpdateBook)
}

func (s *server) HandleDeleteBook() echo.HandlerFunc {
	handler, err := NewDefaultRouteHandler[api.DeleteBookRequestParams, api.DeleteBookResponse](
		WithRequestParamType(api.DeleteBookRequestParams{}),
	)
	if err != nil {
		panic(err)
	}
	return handler(s.controller.DeleteBook)
}
