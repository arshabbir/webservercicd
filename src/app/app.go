package app

import (
	"log"
	"os"

	"github.com/arshabbir/webservercicd/src/controller"

	"github.com/gin-gonic/gin"
)

func StartApplication() {

	version := os.Getenv("VERSION")
	port := os.Getenv("PORT")

	_ = port

	log.Println("Starting the application. Version : ", version)

	bController := controller.NewBookController()

	if bController == nil {
		log.Println("Error creating Book Controller")
		return
	}

	c := gin.Default()

	c.POST("/create", bController.Create)
	c.GET("/read/:id", bController.Read)
	c.PUT("/update", bController.Update)
	c.DELETE("/delete/:id", bController.Delete)

	c.Run(":8080")

	return
}
