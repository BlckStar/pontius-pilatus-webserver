package main

import (
	"strconv"

	"github.com/BlckStar/pontius-pilatus-webserver/console"
	"github.com/BlckStar/pontius-pilatus-webserver/user"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
)

func main() {
	var console = console.Bootstrap()

	var e = echo.New()

	user.Bootstrap(e)

	var port = strconv.Itoa(console.GetPort())

	e.Server.Addr = ":" + port

	e.Logger.Print("Starting Server on Port " + port)
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
