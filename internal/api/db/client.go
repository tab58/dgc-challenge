package db

import "github.com/tab58/dgc-challenge/internal/api/models"

// interface for easier testing
type DBClient interface {
	CreateBook(book models.Book) (int, error)
	GetBook(id int) (models.Book, error)
	GetAllBooks() ([]models.Book, error)
	UpdateBook(id int, book models.Book) (int, error)
	DeleteBook(id int) (bool, error)
}

// concrete type
type someDBClient struct {
	dbConn interface{} // represents a connection to a specific DB instance, ORM, whatever
}

func (c *someDBClient) CreateBook(book models.Book) (int, error) {
	return 0, nil
}

func (c *someDBClient) GetBook(id int) (models.Book, error) {
	return models.Book{}, nil
}

func (c *someDBClient) GetAllBooks() ([]models.Book, error) {
	return []models.Book{}, nil
}

func (c *someDBClient) UpdateBook(id int, book models.Book) (int, error) {
	return id, nil
}

func (c *someDBClient) DeleteBook(id int) (bool, error) {
	return true, nil
}

// This instantiates a new SomeDBClient.
func NewSomeDBClient() (DBClient, error) {
	return &someDBClient{}, nil
}
