package server

import "net/http"

func (s *EchoServer) GetAllCustomers(ctx echo.Context) error {
	email := ctx.QueryParam("email")

	customers, err := s.Db.GetAllCustomers(ctx.Request().Context, email)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError)
	}
}
