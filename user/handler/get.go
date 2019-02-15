package handler

import (
	"net/http"

	"github.com/BlckStar/user"
	"github.com/labstack/echo"
)

func Get(c echo.Context) error {
	var user, err = user.GetUser(c.Param("name"))
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, user)

	return nil
}
