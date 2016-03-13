package main

import (
	"dummy"
	"flag"
	"server"
)

func main() {

	var port int

	flag.IntVar(&port, "Port", 8080, "Port the server listens to")

	flag.Parse()

	go dummy.Run()

	server.Run(uint16(port))

}
