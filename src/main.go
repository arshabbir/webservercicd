package main

import (
	"log"

	"github.com/arshabbir/webservercicd/src/app"
)

func main() {

	log.Println("CICD Webserver  demo.....")
	app.StartApplication()
	return

}
