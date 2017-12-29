package router

import (
	"github.com/labstack/echo"
	"github.com/cappyzawa/go-tiny-crud/api"
)

func NewRouter() *echo.Echo {
	handler := echo.New()
	handler.GET("/users", api.ListUsers)
	handler.GET("/users/:id", api.UserById)
	return handler
}
