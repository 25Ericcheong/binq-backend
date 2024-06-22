package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Binq backend server")

	connStr := "postgres://postgres:LoveS1010_@localhost:8000/binq-pg-dev?sslmode=disable"

	db, dbErr := sql.Open("postgres", connStr)
	if dbErr != nil {
		fmt.Println("Error occurred while trying to setup database")
		log.Fatal(dbErr.Error())
	}

	defer func(db *sql.DB) {
		dbCloseErr := db.Close()
		if dbCloseErr != nil {
			fmt.Println("Error occurred while trying to close database")
			log.Fatal(dbErr.Error())
		}
	}(db)

	dbPingErr := db.Ping()
	if dbPingErr != nil {
		fmt.Println("Error while pinging database")
		log.Fatal(dbPingErr.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Welcome to the index")

		if err != nil {
			return
		}
	})

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println("Error occurred while trying to run server")
		fmt.Println(err.Error())
	}

	fmt.Println("Successful start up. Waiting for request... ")
}
