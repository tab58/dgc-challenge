package models

type RequestParams interface {
	Validate() error
}

type CreateBookRequestParams struct{}

func (p CreateBookRequestParams) Validate() error {
	return nil
}

type GetBookRequestParams struct{}

func (p GetBookRequestParams) Validate() error {
	return nil
}

type UpdateBookRequestParams struct{}

func (p UpdateBookRequestParams) Validate() error {
	return nil
}

type DeleteBookRequestParams struct{}

func (p DeleteBookRequestParams) Validate() error {
	return nil
}

type GetBookListRequestParams struct{}

func (p GetBookListRequestParams) Validate() error {
	return nil
}
