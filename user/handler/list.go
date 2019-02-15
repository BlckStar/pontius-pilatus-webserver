package handler

import (
	"net/http"

	"github.com/BlckStar/user"
	"github.com/labstack/echo"
)

func List(c echo.Context) error {
	c.JSON(http.StatusOK, user.GetAllUsers())

	return nil
}
