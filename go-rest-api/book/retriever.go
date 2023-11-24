package book

import (
	"regexp"

	"github.com/Bash-Clevin/tdd-golang/go-rest-api/rest"
)

type BookFinder interface {
	FindBookBy(isbn string) (rest.Book, error)
}

type Retriever struct {
	f BookFinder
}

func (r Retriever) GetBook(isbn string) (rest.Book, error) {
	b, _ := regexp.MatchString("^[0-9]*$", isbn)

	if !b {
		return rest.Book{}, rest.ErrInvalidISBN
	}

	return r.f.FindBookBy(isbn)
}

func NewRetriever(br BookFinder) Retriever {
	return Retriever{br}
}
