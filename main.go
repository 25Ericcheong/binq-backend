package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// to be able to open postgres driver
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting Binq backend server")

	connStr := "postgres://postgres:LoveS1010_@localhost:8000/binq-pg-dev?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error occurred while trying to setup database")
		log.Fatal(err.Error())
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println("Error occurred while trying to close database")
			log.Fatal(err.Error())
		}
	}(db)

	err = db.Ping()
	if err != nil {
		fmt.Println("Error while pinging database")
		log.Fatal(err.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Welcome to the index")

		if err != nil {
			return
		}
	})

	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println("Error occurred while trying to run server")
		fmt.Println(err.Error())
	}

	fmt.Println("Successful start up. Waiting for request... ")
}
