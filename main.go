package main

import (
	"log"

	"github.com/ramadhanalfarisi/go-grpc/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	app := app.App{}
	app.ConnectDB()
	app.Run()
}