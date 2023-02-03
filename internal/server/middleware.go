package server

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/tab58/dgc-challenge/internal/api"
)

// ValidateParams is a route-level middleware function that validates
// the context parameters in order to determine if the request is valid.
func ValidateParams(params api.RequestParams) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// marshal the context parameters
			err := c.Bind(&params)
			if err != nil {
				// TODO: sending the naked error from bind is probably a bad idea; find a better one
				return api.SendAPIError(c, http.StatusBadRequest, err)
			}

			// check if params are valid
			err = params.Validate()
			if err != nil {
				// TODO: sending the naked error from bind is probably a bad idea; find a better one
				return api.SendAPIError(c, http.StatusBadRequest, err)
			}

			return next(c)
		}
	}
}
