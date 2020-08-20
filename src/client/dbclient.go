package client

import (
	"github.com/arshabbir/webservercicd/src/domain/books_dto"
	"github.com/arshabbir/webservercicd/src/utils"
)

type dbClient struct {
	dbc map[string]books_dto.Book
}

type DBClient interface {
	Create(books_dto.Book) *utils.APIerror
	Read(string) ([]books_dto.Book, *utils.APIerror)
	Update(books_dto.Book) *utils.APIerror
	Delete(string) *utils.APIerror
}

func NewDBClient() DBClient {

	dbc := make(map[string]books_dto.Book, 1)

	return &dbClient{dbc: dbc}

}

func (db *dbClient) Create(b books_dto.Book) *utils.APIerror {

	db.dbc[b.Id] = b

	return nil
}

func (db *dbClient) Read(id string) ([]books_dto.Book, *utils.APIerror) {

	books := make([]books_dto.Book, 0, 1)

	book := db.dbc[id]

	books = append(books, book)

	return books, nil
}

func (db *dbClient) Update(book books_dto.Book) *utils.APIerror {

	db.dbc[book.Id] = book

	return nil
}

func (db *dbClient) Delete(id string) *utils.APIerror {

	delete(db.dbc, id)

	return nil
}
