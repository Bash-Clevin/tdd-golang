package e2etest_test

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type GetSingleBookSuite struct {
	suite.Suite
}

func TestGetSingleBookSuite(t *testing.T) {
	suite.Run(t, new(GetSingleBookSuite))
}

func (s *GetSingleBookSuite) TestGetBookThatDoesNotExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/123456789")
	body, _ := io.ReadAll(r.Body)

	s.Equal(http.StatusNotFound, r.StatusCode)
	s.JSONEq(`{"code": "001", "msg": "No Book with ISBN 123456789"}`, string(body))
}

func (s *GetSingleBookSuite) TestGetBookWithInvalidISBN() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/1234C6789")
	body, _ := io.ReadAll(r.Body)

	s.Equal(http.StatusBadRequest, r.StatusCode)
	s.JSONEq(`{"code": "003", "msg": "ISBN is invalid"}`, string(body))
}

func (s *GetSingleBookSuite) TestGetBookThatDoesExist() {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	db.Exec("INSERT INTO book (isbn, name,  image, genre, year_published) VALUES ('987654321', 'Testing all stuff', 'testing.jpg', 'computing', 2021 )")

	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/book/987654321")
	body, _ := io.ReadAll(r.Body)

	s.Equal(http.StatusOK, r.StatusCode)

	expectedBody := `{
		"isbn": "987654321",
		"name": "Testing all stuff",
		"image": "testing.jpg",
		"genre": "computing",
		"year_published": 2021
	}`
	s.JSONEq(expectedBody, string(body))

}
