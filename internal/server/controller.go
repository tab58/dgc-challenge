package server

import (
	"fmt"

	"github.com/tab58/dgc-challenge/internal/api"
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
}

func (c *defaultController) GetBookList(params api.GetBookListRequestParams) (api.GetBookListResponse, error) {
	resp := api.GetBookListResponse{}

	return resp, fmt.Errorf("not implemented in controller")
}

func (c *defaultController) CreateBook(params api.CreateBookRequestParams) (api.CreateBookResponse, error) {
	resp := api.CreateBookResponse{}

	return resp, fmt.Errorf("not implemented in controller")
}

func (c *defaultController) GetBook(params api.GetBookRequestParams) (api.GetBookResponse, error) {
	resp := api.GetBookResponse{}

	return resp, fmt.Errorf("not implemented in controller")
}

func (c *defaultController) UpdateBook(params api.UpdateBookRequestParams) (api.UpdateBookResponse, error) {
	resp := api.UpdateBookResponse{}

	return resp, fmt.Errorf("not implemented in controller")
}

func (c *defaultController) DeleteBook(params api.DeleteBookRequestParams) (api.DeleteBookResponse, error) {
	resp := api.DeleteBookResponse{}

	return resp, fmt.Errorf("not implemented in controller")
}

func NewController() (Controller, error) {
	controller := &defaultController{}
	return controller, nil
}
