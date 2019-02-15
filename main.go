package main

import (
	"github.com/BlckStar/pontius-pilatus-webserver/user"
	"github.com/labstack/echo"
)

func main() {
	var e = echo.New()

	user.Bootstrap(e)

	e.Logger.Fatal(e.Start(":1323"))
}
