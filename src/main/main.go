package main

import (
	"flag"
	"server"
)

func main() {

	var port int

	flag.IntVar(&port, "Port", 8080, "Port the server listens to")

	flag.Parse()
	server.Run(port)

}
