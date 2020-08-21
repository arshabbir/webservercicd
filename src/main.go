package main

import (
	"log"

	"github.com/arshabbir/webservercicd/src/app"
)

func main() {

	log.Println("CICD Webserve demo.....")
	app.StartApplication()
	return

}
