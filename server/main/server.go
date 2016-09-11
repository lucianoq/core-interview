package main

import (
	"core-interview/server/webserver"
	"flag"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Set a port. Default 8080")
	flag.Parse()
	webserver.Start(port)
}
