package server

import (
	"fmt"

	"github.com/tab58/dgc-challenge/internal/api"
	"github.com/tab58/dgc-challenge/internal/api/db"
	"github.com/tab58/dgc-challenge/internal/api/models"
)

type Controller interface {
	GetBookList(params api.GetBookListRequestParams) (api.GetBookListResponse, error)
	CreateBook(params api.CreateBookRequestParams) (api.CreateBookResponse, error)
	GetBook(params api.GetBookRequestParams) (api.GetBookResponse, error)
	UpdateBook(params api.UpdateBookRequestParams) (api.UpdateBookResponse, error)
	DeleteBook(params api.DeleteBookRequestParams) (api.DeleteBookResponse, error)
}

type defaultController struct {
	// add dependencies here
	dbClient db.DBClient
}

func (c *defaultController) GetBookList(params api.GetBookListRequestParams) (api.GetBookListResponse, error) {
	books, err := c.dbClient.GetAllBooks()
	if err != nil {
		return api.GetBookListResponse{}, err
	}

	resp := api.GetBookListResponse{
		Count:    len(books),
		Next:     "null", // this is just for pagination, but I got stuck here
		Previous: "null", // this is just for pagination, but I got stuck here
		Results:  books,
	}
	return resp, nil
}

func (c *defaultController) CreateBook(params api.CreateBookRequestParams) (api.CreateBookResponse, error) {
	book := models.Book{
		ISBN: params.ISBN,
	}

	id, err := c.dbClient.CreateBook(book)
	if err != nil {
		return api.CreateBookResponse{}, err
	}

	resp := api.CreateBookResponse{
		ID:   id,
		Data: book,
	}
	return resp, nil
}

func (c *defaultController) GetBook(params api.GetBookRequestParams) (api.GetBookResponse, error) {
	book, err := c.dbClient.GetBook(params.ID)
	if err != nil {
		return api.GetBookResponse{}, err
	}

	resp := api.GetBookResponse{
		Data: book,
	}
	return resp, nil
}

func (c *defaultController) UpdateBook(params api.UpdateBookRequestParams) (api.UpdateBookResponse, error) {
	updatedBook := models.Book{
		ISBN: params.ISBN,
	}

	id, err := c.dbClient.UpdateBook(params.ID, updatedBook)
	if err != nil {
		return api.UpdateBookResponse{}, err
	}

	resp := api.UpdateBookResponse{
		ID:   id,
		Data: updatedBook,
	}
	return resp, nil
}

func (c *defaultController) DeleteBook(params api.DeleteBookRequestParams) (api.DeleteBookResponse, error) {
	id := params.ID
	isDeleted, err := c.dbClient.DeleteBook(id)
	if err != nil {
		return api.DeleteBookResponse{}, err
	}
	if !isDeleted {
		return api.DeleteBookResponse{}, fmt.Errorf("unable to delete entity Book")
	}

	resp := api.DeleteBookResponse{
		ID: id,
	}
	return resp, nil
}

func NewController(dbClient db.DBClient) (Controller, error) {
	controller := &defaultController{
		dbClient: dbClient,
	}
	return controller, nil
}
