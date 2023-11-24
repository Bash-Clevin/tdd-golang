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

type DBRetrieverSuite struct {
	suite.Suite
}

func TestDBRetrieverSuite(t *testing.T) {
	suite.Run(t, new(DBRetrieverSuite))
}

var (
	db *sql.DB
	r  book.DBRetriever
)

func (s *DBRetrieverSuite) SetupTest() {
	db, _ = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	r = book.NewDBRetriever(db)
}

func (s *DBRetrieverSuite) TearDownTest() {
	db.Close()
}

func (s *DBRetrieverSuite) TestRetrievBookThatDoesNotExist() {
	_, err := r.FindBookBy("123456789")

	s.Equal(rest.ErrBookNotFound, err)
}

func (s *DBRetrieverSuite) TestGetBookThatDoesExist() {
	db.Exec("INSERT INTO book (isbn, name,  image, genre, year_published) VALUES ('987654321', 'Testing all stuff', 'testing.jpg', 'computing', 2021 )")

	b, err := r.FindBookBy("987654321")

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

func (s *DBRetrieverSuite) TestUnexpectedErrorRetrievingBook() {
	db.Close()

	_, err := r.FindBookBy("123456789")

	s.Equal(book.ErrFailedToRetrieve, err)
}
