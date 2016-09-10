package main

import (
	"core-interview/server/webserver"
	"core-interview/server/storage"
	"log"
	"os"
)

func main() {
	err := storage.Check()
	if err != nil {
		log.Print(err.Error())
		os.Exit(-1)
	}
	defer storage.DBClose()

	webserver.Start()
}
