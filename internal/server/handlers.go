package server

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/tab58/dgc-challenge/internal/api"

	"reflect"
)

type routeHandlerData struct {
	successHTTPCode int
	failureHTTPCode int
	reqParamType    api.RequestParams
}

type RouteHandlerOption func(*routeHandlerData)

// WithSuccessCode changes the returned HTTP code on success.
func WithSuccessCode(code int) func(*routeHandlerData) {
	return func(r *routeHandlerData) {
		r.successHTTPCode = code
	}
}

// WithFailureCode changes the returned HTTP code on failure.
func WithFailureCode(code int) func(*routeHandlerData) {
	return func(r *routeHandlerData) {
		r.failureHTTPCode = code
	}
}

// WithRequestParamType takes in a value and creates a reflect.Type so that the route handler can get a new struct from that type when the function is run.
func WithRequestParamType(valueForType api.RequestParams) func(*routeHandlerData) {
	return func(r *routeHandlerData) {
		r.reqParamType = valueForType
	}
}

// NewDefaultRouteHandler creates an echo.HandlerFunc that embeds the logic to handle the success
// and failure cases for a controller function that maps api.RequestParams input to a response/error output.
func NewDefaultRouteHandler[T api.RequestParams, U interface{}](options ...RouteHandlerOption) (func(handler func(params T) (U, error)) echo.HandlerFunc, error) {
	data := &routeHandlerData{
		successHTTPCode: http.StatusOK,
		failureHTTPCode: http.StatusInternalServerError,
		reqParamType:    nil,
	}

	// apply all options
	for _, option := range options {
		option(data)
	}

	// check for required option values
	if data.reqParamType == nil {
		return nil, fmt.Errorf("route handler must have request parameter type defined")
	}

	// create the handler decorator
	h := func(handler func(T) (U, error)) echo.HandlerFunc {
		paramType := reflect.TypeOf(data.reqParamType)

		// create the actual route handler function
		return func(ctx echo.Context) error {
			// bind the request parameters
			params := reflect.New(paramType).Elem().Interface().(T)
			err := ctx.Bind(&params)
			if err != nil {
				return api.SendAPIError(ctx, data.failureHTTPCode, err)
			}

			// run the function argument
			resp, err := handler(params)

			// handle success/failure cases
			if err != nil {
				return api.SendAPIError(ctx, data.failureHTTPCode, err)
			}
			return ctx.JSON(data.successHTTPCode, resp)
		}
	}

	return h, nil
}
