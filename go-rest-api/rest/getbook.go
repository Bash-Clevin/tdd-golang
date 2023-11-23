package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ISBN          string `json:"isbn"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	Genre         string `json:"genre"`
	YearPublished int    `json:"year_published"`
}

type jsonError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

var (
	ErrBookNotFound = errors.New("Book not found")
)

type BookRetriver interface {
	GetBook(isbn string) (Book, error)
}

type GetBookHandler struct {
	br BookRetriver
}

func (g GetBookHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	isbn := v["isbn"]

	book, err := g.br.GetBook(isbn)

	if err != nil {
		var e jsonError
		if err == ErrBookNotFound {
			e.Code = "001"
			e.Msg = fmt.Sprintf("No Book with ISBN %s", isbn)

			w.WriteHeader(http.StatusNotFound)
		} else {
			e.Code = "002"
			e.Msg = "Error attempting to get book"

			w.WriteHeader(http.StatusInternalServerError)
		}

		body, _ := json.Marshal(e)

		w.Write(body)

		return
	}

	body, _ := json.Marshal(book)

	w.WriteHeader(http.StatusOK)

	w.Write(body)

}

func NewGetBookHandler(br BookRetriver) GetBookHandler {
	return GetBookHandler{br}
}
