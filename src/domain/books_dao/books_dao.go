package books_dao

import (
	"github.com/arshabbir/webservercicd/src/client"
	"github.com/arshabbir/webservercicd/src/domain/books_dto"
	"github.com/arshabbir/webservercicd/src/utils"
)

type bookdDao struct {
	dbClient client.DBClient
}

type BooksDao interface {
	Create(books_dto.Book) *utils.APIerror
	Read(string) ([]books_dto.Book, *utils.APIerror)
	Update(books_dto.Book) *utils.APIerror
	Delete(string) *utils.APIerror
}

func NewDAO() BooksDao {

	dbClient := client.NewDBClient()

	return &bookdDao{dbClient: dbClient}
}

func (bd *bookdDao) Create(b books_dto.Book) *utils.APIerror {

	return bd.dbClient.Create(b)
}

func (bd *bookdDao) Read(id string) ([]books_dto.Book, *utils.APIerror) {
	return bd.dbClient.Read(id)
}

func (bd *bookdDao) Update(book books_dto.Book) *utils.APIerror {

	return bd.dbClient.Update(book)
}

func (bd *bookdDao) Delete(id string) *utils.APIerror {
	return bd.dbClient.Delete(id)
}
