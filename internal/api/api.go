package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tab58/dgc-challenge/internal/api/models"
)

type APIErrorResponse struct {
	Message string `json:"message"`
}

func SendAPIError(c echo.Context, code int, err error) error {
	return c.JSON(code, APIErrorResponse{Message: err.Error()})
}

// RequestParams is an interface for any type of HTTP request parameters.
type RequestParams interface {
	Validate() error
}

// Request parameters for the book list page (aggregate).
type GetBookListRequestParams struct{}

func (p GetBookListRequestParams) Validate() error {
	// check bound parameters here
	return nil
}

type GetBookListResponse struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []models.Book `json:"results"`
}

// Request parameters for creating a book.
type CreateBookRequestParams struct{}

func (p CreateBookRequestParams) Validate() error {
	// check bound parameters here
	return nil
}

type CreateBookResponse struct {
	ID   int         `json:"id"`
	Data models.Book `json:"data"`
}

// Request parameters for retrieving a book.
type GetBookRequestParams struct {
	ID int `param:"id"`
}

func (p GetBookRequestParams) Validate() error {
	// check bound parameters here
	return nil
}

type GetBookResponse struct {
	Data models.Book `json:"data"`
}

// Request parameters for updating a book.
type UpdateBookRequestParams struct {
	ID int `param:"id"`
}

type UpdateBookResponse struct {
	ID   int         `json:"id"`
	Data models.Book `json:"data"`
}

func (p UpdateBookRequestParams) Validate() error {
	// check bound parameters here
	return nil
}

// Request parameters for deleting a book.
type DeleteBookRequestParams struct {
	ID int `param:"id"`
}

func (p DeleteBookRequestParams) Validate() error {
	// check bound parameters here
	return nil
}

type DeleteBookResponse struct {
	ID int `json:"id"`
}
