package api

import (
	"net/http"
	"github.com/labstack/echo"
)

var (
	user  User
	users []User
)

func ListUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, GetUsers(users))
}

func UserById(c echo.Context) error{
	return c.JSON(http.StatusOK, GetUser(user, c.Param("id")))
}
