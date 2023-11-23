package book

import (
	"database/sql"
	"errors"

	"github.com/Bash-Clevin/tdd-golang/go-rest-api/rest"
)

type Retriever struct {
	db *sql.DB
}

var (
	ErrFailedToRetrieve = errors.New("error occured when retrieving book")
)

func (r Retriever) GetBook(isbn string) (rest.Book, error) {
	b := rest.Book{}

	row := r.db.QueryRow("SELECT isbn, name, image, genre, year_published FROM book WHERE isbn = $1", isbn)
	err := row.Scan(
		&b.ISBN,
		&b.Name,
		&b.Image,
		&b.Genre,
		&b.YearPublished,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return b, rest.ErrBookNotFound
		}
		return b, ErrFailedToRetrieve
	}

	return b, nil
}

func NewRetriever(db *sql.DB) Retriever {
	return Retriever{db}
}
