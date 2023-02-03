package server

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/tab58/dgc-challenge/internal/api/models"
)

func ValidateParams(params models.RequestParams) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// marshal the context parameters
			err := c.Bind(&params)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}

			// check if params are valid
			paramErr := params.Validate()
			if paramErr != nil {
				return paramErr
			}

			return next(c)
		}
	}
}
