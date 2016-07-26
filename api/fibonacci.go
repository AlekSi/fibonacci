package api

import (
	"strconv"

	"github.com/AlekSi/fibonacci/errs"
	"github.com/labstack/echo"
)

// GetFibonacci GET /fibonacci/:n - returns Fibonacci numbers.
func (h *Handler) GetFibonacci(ctx echo.Context) error {
	n, err := strconv.ParseUint(ctx.Param("n"), 10, 64)
	if n == 0 {
		return errs.New(errs.InvalidParameter, "'n' is not a natural number", err)
	}
	numbers, err := h.f.Numbers(uint(n))
	if err != nil {
		return err
	}
	return ctx.JSON(200, numbers)
}
