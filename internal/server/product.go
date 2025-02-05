package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	products, err := s.DB.GetAllProducts(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, products)
}
