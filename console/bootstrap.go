package console

import (
	"flag"
)

type Console struct {
	port *int
}

func Bootstrap() Console {
	var port = flag.Int("port", 1121, "the Port the Application should listen to")

	flag.Parse()

	return Console{
		port: port,
	}
}

func (c *Console) GetPort() int {
	return *c.port
}
