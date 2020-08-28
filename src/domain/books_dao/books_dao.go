package books_dao

import (
	"fmt"
	"log"
	"os"

	"github.com/arshabbir/webservercicd/src/domain/books_dto"
	"github.com/arshabbir/webservercicd/src/utils"
	"github.com/gocql/gocql"
)

type bookDao struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

type BooksDao interface {
	Create(books_dto.Book) *utils.APIerror
	Read(string) ([]books_dto.Book, *utils.APIerror)
	Update(books_dto.Book) *utils.APIerror
	Delete(string) *utils.APIerror
}

func NewDAO() BooksDao {

	//Get the Environment variable "CASSANDRACLUSTER"

	clusterIP := os.Getenv("CLUSTERIP")

	log.Println("ClusterIP environment  : ", clusterIP)

	cluster := gocql.NewCluster(clusterIP)
	cluster.Keyspace = "bookstore"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()

	if err != nil {
		log.Println("Cassandra Session Creation Error..", err.Error())
		return nil
	}

	//defer session.Close()
	return &bookDao{cluster: cluster, session: session}

}

func (bd *bookDao) Create(b books_dto.Book) *utils.APIerror {

	log.Println("Creating the book...")

	/*
				Id          string `json:"id"`
			Name        string `json:"name"`
			Author      string `json:"author"`
			Publication string `json:"publication"`
			Pages       int    `json:"pages"`
		}

	*/

	insertQuery := fmt.Sprintf("INSERT INTO books(id, name, author, publication, pages) VALUES(?, ?, ?, ?,?);")

	err := bd.session.Query(insertQuery, b.Id, b.Name, b.Author, b.Publication, b.Pages).Consistency(gocql.Quorum).Exec()

	if err != nil {
		log.Println("Errror in insertion")
		return &utils.APIerror{Status: 500, Msg: "Error in Insertion"}
	}
	return nil
}

func (bd *bookDao) Read(id string) ([]books_dto.Book, *utils.APIerror) {
	log.Println("Reading the book...")

	return nil, nil
}

func (bd *bookDao) Update(book books_dto.Book) *utils.APIerror {

	log.Println("Updating the book...")

	return nil
}

func (bd *bookDao) Delete(id string) *utils.APIerror {
	log.Println("Deleting the book...")

	return nil
}
