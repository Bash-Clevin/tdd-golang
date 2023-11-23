package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"

	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type health struct {
	Status   string   `json:"status"`
	Messages []string `json:"messages"`
}

type jsonError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type Book struct {
	ISBN          string `json:"isbn"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	Genre         string `json:"genre"`
	YearPublished int    `json:"year_published"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to DB: %s", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatalf("Error making DB driver: %s", err.Error())
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	migrator.Steps(2)

	if err != nil {
		log.Fatalf("Error making migration engine: %s", err.Error())
	}

	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		h := health{
			Status:   "OK",
			Messages: []string{},
		}

		b, _ := json.Marshal(h)

		w.Write(b)

		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/book/{isbn}", func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)

		isbn := v["isbn"]
		e := jsonError{
			Code: "001",
			Msg:  fmt.Sprintf("No Book with ISBN %s", isbn),
		}

		b := Book{}
		row := db.QueryRow("SELECT isbn, name, image, genre, year_published FROM book WHERE isbn = $1", isbn)
		err := row.Scan(
			&b.ISBN,
			&b.Name,
			&b.Image,
			&b.Genre,
			&b.YearPublished,
		)

		if err != nil {
			body, _ := json.Marshal(e)

			w.WriteHeader(http.StatusNotFound)

			w.Write(body)

			return
		}

		body, _ := json.Marshal(b)

		w.WriteHeader(http.StatusOK)

		w.Write(body)

	})

	s := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
