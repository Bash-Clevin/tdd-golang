package book

import (
	"database/sql"
	"errors"

	"github.com/Bash-Clevin/tdd-golang/go-rest-api/rest"
)

type DBRetriever struct {
	db *sql.DB
}

var (
	ErrFailedToRetrieve = errors.New("error occured when retrieving book")
)

func (r DBRetriever) FindBookBy(isbn string) (rest.Book, error) {
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

func NewDBRetriever(db *sql.DB) DBRetriever {
	return DBRetriever{db}
}
