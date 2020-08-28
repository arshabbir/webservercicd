package services

import (
	"log"

	"github.com/arshabbir/webservercicd/src/domain/books_dao"
	"github.com/arshabbir/webservercicd/src/domain/books_dto"
	"github.com/arshabbir/webservercicd/src/utils"
)

type bookService struct {
	bdao books_dao.BooksDao
}

type BookService interface {
	Create(books_dto.Book) *utils.APIerror
	Read(string) ([]books_dto.Book, *utils.APIerror)
	Update(books_dto.Book) *utils.APIerror
	Delete(string) *utils.APIerror
}

func NewBookService() BookService {

	bdao := books_dao.NewDAO()

	if bdao == nil {
		log.Println("Error createing the DAO object")
		return nil
	}

	log.Println("DAO obhect creation successful")

	return &bookService{bdao: bdao}

}

func (bs *bookService) Create(b books_dto.Book) *utils.APIerror {

	return bs.bdao.Create(b)
}

func (bs *bookService) Read(id string) ([]books_dto.Book, *utils.APIerror) {
	return bs.bdao.Read(id)
}

func (bs *bookService) Update(book books_dto.Book) *utils.APIerror {

	return bs.bdao.Update(book)
}

func (bs *bookService) Delete(id string) *utils.APIerror {
	return bs.bdao.Delete(id)
}
