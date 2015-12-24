package main

import (
	"flag"
	"server"
)

func main() {

	var port int

	flag.IntVar(&port, "Port", 8080, "Port the server listens to")

	flag.Parse()

	game.Open()
	defer game.Close()
	server.Run(port)

}
