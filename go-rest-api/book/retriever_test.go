package book_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/Bash-Clevin/tdd-golang/go-rest-api/book"
	"github.com/Bash-Clevin/tdd-golang/go-rest-api/rest"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type RetrieverSuite struct {
	suite.Suite
}

func TestRetrieverSuite(t *testing.T) {
	suite.Run(t, new(RetrieverSuite))
}

var (
	db *sql.DB
	r  book.Retriever
)

func (s *RetrieverSuite) SetupTest() {
	db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	r = book.NewRetriever(db)
}

func (s *RetrieverSuite) TearDownTest() {
	db.Close()
}

func (s *RetrieverSuite) TestRetrievBookThatDoesNotExist() {
	_, err := r.GetBook("123456789")

	s.Equal(rest.ErrBookNotFound, err)
}

func (s *RetrieverSuite) TestGetBookThatDoesExist() {
	db.Exec("INSERT INTO book (isbn, name,  image, genre, year_published) VALUES ('987654321', 'Testing all stuff', 'testing.jpg', 'computing', 2021 )")

	b, err := r.GetBook("987654321")

	s.NoError(err)

	book := rest.Book{
		ISBN:          "987654321",
		Name:          "Testing all stuff",
		Image:         "testing.jpg",
		Genre:         "computing",
		YearPublished: 2021,
	}

	s.Equal(book, b)

}

func (s *RetrieverSuite) TestUnexpectedErrorRetrievingBook() {
	db.Close()

	_, err := r.GetBook("123456789")

	s.Equal(book.ErrFailedToRetrieve, err)
}
