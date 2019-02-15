package user

import (
	"github.com/BlckStar/pontius-pilatus-webserver/user/handler"
	"github.com/labstack/echo"
)

func Bootstrap(e *echo.Echo) {
	e.GET("/users", handler.List)
	e.GET("/users/:name", handler.Get)
}
