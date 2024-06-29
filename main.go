package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	// to be able to open postgres driver
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting Binq backend server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DB_CONNECTION_STRING")

	db, err := sql.Open(os.Getenv("DB_DRIVER"), connStr)
	if err != nil {
		fmt.Println("Error occurred while trying to setup database")
		log.Fatal(err.Error())
	}

	createTicketTable(db)

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

	err = http.ListenAndServe(os.Getenv("URL"), mux)
	if err != nil {
		fmt.Println("Error occurred while trying to run server")
		fmt.Println(err.Error())
	}

	fmt.Println("Successful start up. Waiting for request... ")
}

// Real application uses migration
func createTicketTable(db *sql.DB) {
	/* Ticket Table
	id SERIAL - so PK auto create
	ticket_num
	branch
	customer_name
	customer_pax_num
	customer_phone
	created_on_date_time DEFAULT NOW() - so by default inserts today's date
	*/
	query := "CREATE TABLE IF NOT EXISTS ticket" +
		"(" +
		"id SERIAL PRIMARY KEY," +
		"ticket_num INTEGER NOT NULL CHECK(ticket_num > -1)," +
		"branch VARCHAR(50) NOT NULL," +
		"customer_name VARCHAR(100) NOT NULL," +
		"customer_pax_num INTEGER NOT NULL CHECK(customer_pax_num > 0)," +
		"customer_phone VARCHAR(20) NOT NULL," +
		"created_on_date_time TIMESTAMP NOT NULL DEFAULT NOW()" +
		")"

	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error while creating ticket table")
		log.Fatal(err.Error())
	}
}
