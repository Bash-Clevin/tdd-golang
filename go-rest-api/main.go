package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Bash-Clevin/tdd-golang/go-rest-api/book"
	"github.com/Bash-Clevin/tdd-golang/go-rest-api/rest"
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

	retriever := book.NewRetriever(db)

	r.Handle("/book/{isbn}", rest.NewGetBookHandler(retriever))

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		h := health{
			Status:   "OK",
			Messages: []string{},
		}

		b, _ := json.Marshal(h)

		w.WriteHeader(http.StatusOK)

		w.Write(b)
	})

	s := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
