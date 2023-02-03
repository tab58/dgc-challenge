package server

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/tab58/dgc-challenge/internal/api/models"
)

func AttachRoutes(s *server) {
	s.srv.GET("/books",
		s.HandleGetBookList,
		ValidateParams(models.GetBookListRequestParams{}),
	)

	s.srv.POST("/books",
		s.HandleCreateBook,
		ValidateParams(models.GetBookRequestParams{}),
	)
	s.srv.GET("/books/:id",
		s.HandleGetBook,
		ValidateParams(models.GetBookRequestParams{}),
	)
	s.srv.PUT("/books/:id",
		s.HandleUpdateBook,
		ValidateParams(models.UpdateBookRequestParams{}),
	)
	s.srv.DELETE("/books/:id",
		s.HandleDeleteBook,
		ValidateParams(models.DeleteBookRequestParams{}),
	)
}

func (s *server) HandleGetBookList(ctx echo.Context) error {
	return fmt.Errorf("not implemented")
}

func (s *server) HandleCreateBook(ctx echo.Context) error {
	return fmt.Errorf("not implemented")
}

func (s *server) HandleGetBook(ctx echo.Context) error {
	return fmt.Errorf("not implemented")
}

func (s *server) HandleUpdateBook(ctx echo.Context) error {
	return fmt.Errorf("not implemented")
}

func (s *server) HandleDeleteBook(ctx echo.Context) error {
	return fmt.Errorf("not implemented")
}
