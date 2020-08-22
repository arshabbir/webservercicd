package controller

import (
	"log"
	"net/http"

	"github.com/arshabbir/webservercicd/src/domain/books_dto"
	"github.com/arshabbir/webservercicd/src/services"
	"github.com/arshabbir/webservercicd/src/utils"
	"github.com/gin-gonic/gin"
)

type bookController struct {
	service services.BookService
}

type BookController interface {
	Create(*gin.Context)
	Read(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	Ping(*gin.Context)
	Wish(*gin.Context)
}

func NewBookController() BookController {

	service := services.NewBookService()

	return &bookController{service: service}
}

func (bc *bookController) Create(c *gin.Context) {

	var book books_dto.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, &utils.APIerror{Status: http.StatusBadRequest, Msg: "Bad request"})
		return
	}

	if err := bc.service.Create(book); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.APIerror{Status: http.StatusInternalServerError, Msg: "Book Creation Error"})
		return
	}

	c.JSON(http.StatusInternalServerError, &utils.APIerror{Status: http.StatusOK, Msg: "Book Creation Successfull"})
	return

}

func (bc *bookController) Read(c *gin.Context) {

	id := c.Param("id")

	book, err := bc.service.Read(id)

	if err != nil {
		c.JSON(http.StatusNotFound, &utils.APIerror{Status: http.StatusNotFound, Msg: "No book found"})
		return
	}
	c.JSON(http.StatusOK, &book)

	return
}

func (bc *bookController) Update(c *gin.Context) {

	var book books_dto.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, &utils.APIerror{Status: http.StatusBadRequest, Msg: "Bad request"})
		return
	}

	if err := bc.service.Update(book); err != nil {
		c.JSON(http.StatusInternalServerError, &utils.APIerror{Status: http.StatusInternalServerError, Msg: "Book Update Error"})
		return
	}

	c.JSON(http.StatusInternalServerError, &utils.APIerror{Status: http.StatusOK, Msg: "Book Update Successfull"})
	return
}

func (bc *bookController) Delete(c *gin.Context) {
	id := c.Param("id")

	err := bc.service.Delete(id)

	if err != nil {
		c.JSON(http.StatusNotFound, &utils.APIerror{Status: http.StatusNotFound, Msg: "No book found"})
		return
	}
	c.JSON(http.StatusOK, &utils.APIerror{Status: http.StatusNotFound, Msg: "Book deletion successfull"})

	return
}

func (bc *bookController) Ping(c *gin.Context) {

	c.String(http.StatusOK, "pongv2.0")

	return
}

func (bc *bookController) Wish(c *gin.Context) {

	log.Println("Wish API invoked")

	c.File(".\\controller\\flowers.jpg")
	return
}
