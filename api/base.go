package api

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/AlekSi/fibonacci/errs"
	"github.com/AlekSi/fibonacci/services"
)

// Handler stores request-independent state.
type Handler struct {
	f *services.Fibonacci
}

// NewHandler creates new handler.
func NewHandler() *Handler {
	return &Handler{
		f: services.NewFibonacci(),
	}
}

// ErrorHandler is an Echo error handler for our errors and API conventions.
func ErrorHandler(err error, ctx echo.Context) {
	ctx.Logger().Errorf("ErrorHandler: %v (%T)", err, err)

	switch err := err.(type) {
	case *errs.APIError:
		// expose original error only in debug mode
		if !ctx.Echo().Debug() {
			err.Err = nil
		}
		ctx.JSON(int(err.Code)/1000, err)

	// invalid route, etc.
	case *echo.HTTPError:
		ctx.String(err.Code, err.Message)

	default:
		if ctx.Echo().Debug() {
			ctx.String(500, fmt.Sprintf("Internal Server Error: %s", err))
		} else {
			ctx.String(500, "Internal Server Error")
		}
	}
}
