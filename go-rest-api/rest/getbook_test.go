package rest_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bash-Clevin/tdd-golang/go-rest-api/rest"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GetBookSuite struct {
	suite.Suite
}

func TestGetBookSuite(t *testing.T) {
	suite.Run(t, new(GetBookSuite))
}

func (s *GetBookSuite) TestGetBookThatDoesNotExist() {
	req, _ := http.NewRequest(http.MethodGet, "/book/123456789", nil)
	req = mux.SetURLVars(req, map[string]string{"isbn": "123456789"})
	resp := httptest.NewRecorder()

	br := new(MockBookRetriever)
	br.On("GetBook", "123456789").Return(rest.Book{}, rest.ErrBookNotFound)

	h := rest.NewGetBookHandler(br)

	h.ServeHttp(resp, req)

	body, _ := io.ReadAll(resp.Body)

	s.Equal(http.StatusNotFound, resp.Code)
	s.JSONEq(`{"code": "001", "msg": "No Book with ISBN 123456789"}`, string(body))
}

func (s *GetBookSuite) TestGetBookThatDoesExist() {
	req, _ := http.NewRequest(http.MethodGet, "/book/123456789", nil)
	req = mux.SetURLVars(req, map[string]string{"isbn": "123456789"})
	resp := httptest.NewRecorder()

	br := new(MockBookRetriever)
	book := rest.Book{
		ISBN:          "987654321",
		Name:          "Testing all stuff",
		Image:         "testing.jpg",
		Genre:         "computing",
		YearPublished: 2021,
	}
	br.On("GetBook", "123456789").Return(book, nil)

	h := rest.NewGetBookHandler(br)

	h.ServeHttp(resp, req)

	body, _ := io.ReadAll(resp.Body)

	s.Equal(http.StatusOK, resp.Code)

	expectedBody := `{
		"isbn": "987654321",
		"name": "Testing all stuff",
		"image": "testing.jpg",
		"genre": "computing",
		"year_published": 2021
	}`

	s.JSONEq(expectedBody, string(body))
}

func (s *GetBookSuite) TestBookReturnUnexpectedError() {
	req, _ := http.NewRequest(http.MethodGet, "/book/123456789", nil)
	req = mux.SetURLVars(req, map[string]string{"isbn": "123456789"})
	resp := httptest.NewRecorder()

	br := new(MockBookRetriever)
	br.On("GetBook", "123456789").Return(rest.Book{}, errors.New("unexpected error"))

	h := rest.NewGetBookHandler(br)

	h.ServeHttp(resp, req)

	body, _ := io.ReadAll(resp.Body)

	s.Equal(http.StatusInternalServerError, resp.Code)
	s.JSONEq(`{"code": "002", "msg": "Error attempting to get book"}`, string(body))
}

type MockBookRetriever struct {
	mock.Mock
}

func (m *MockBookRetriever) GetBook(isbn string) (rest.Book, error) {
	args := m.Called(isbn)

	return args.Get(0).(rest.Book), args.Error(1)
}
