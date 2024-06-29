package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	// to be able to open postgres driver
	_ "github.com/lib/pq"
)

type Ticket struct {
	Branch         string
	CustomerName   string
	CustomerPaxNum int
	CustomerPhone  string
}

type DbTicket struct {
	TicketId          string
	Branch            string
	CustomerName      string
	CustomerPaxNum    int
	CustomerPhone     string
	CreatedOnDateTime time.Time
}

//"branch VARCHAR(50) NOT NULL," +
//"customer_name VARCHAR(100) NOT NULL," +
//"customer_pax_num INTEGER NOT NULL CHECK(customer_pax_num > 0)," +
//"customer_phone VARCHAR(20) NOT NULL," +

func main() {
	time.Now()
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

	newTicket := Ticket{"SS15", "Eric", 5, "0122817216"}
	rowTicket := insertDbTicket(db, newTicket)

	fmt.Printf("Ticket details \n"+
		"Id: %s \n"+
		"Customer Name: %s \n"+
		"Creation Date: %s \n", rowTicket.TicketId, rowTicket.CustomerName, rowTicket.CreatedOnDateTime)

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

	url := os.Getenv("URL")
	fmt.Println("Successful! Listening on: " + url)
	err = http.ListenAndServe(url, mux)
	if err != nil {
		fmt.Println("Error occurred while trying to run server")
		fmt.Println(err.Error())
	}
}

// Real application uses migration
// TODO: Setup migration setup during deployment as well as testing
func createTicketTable(db *sql.DB) {
	/* Ticket Table
	id SERIAL - so PK auto create
	ticket_id - used to display to customer
	branch
	customer_name
	customer_pax_num
	customer_phone
	created_on_date_time DEFAULT NOW() - so by default inserts today's date
	*/
	query := "CREATE TABLE IF NOT EXISTS ticket" +
		"(id SERIAL PRIMARY KEY," +
		"ticket_id VARCHAR(20) UNIQUE NOT NULL," +
		"branch VARCHAR(50) NOT NULL," +
		"customer_name VARCHAR(100) NOT NULL," +
		"customer_pax_num INTEGER NOT NULL CHECK(customer_pax_num > 0)," +
		"customer_phone VARCHAR(20) NOT NULL," +
		"created_on_date_time TIMESTAMP NOT NULL DEFAULT NOW())"

	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error while creating ticket table")
		log.Fatal(err.Error())
	}
}

func insertDbTicket(db *sql.DB, ticket Ticket) (row DbTicket) {
	query := "INSERT INTO ticket" +
		"(ticket_id, branch, customer_name, customer_pax_num, customer_phone) " +
		"VALUES ($1, $2, $3, $4, $5)" +
		"RETURNING ticket_id, branch, customer_name, customer_pax_num, customer_phone, created_on_date_time"

	ticketId := ticket.Branch + "_1"

	err := db.
		QueryRow(query, ticketId, ticket.Branch, ticket.CustomerName, ticket.CustomerPaxNum, ticket.CustomerPhone).
		Scan(&row.TicketId, &row.Branch, &row.CustomerName, &row.CustomerPaxNum, &row.CustomerPhone, &row.CreatedOnDateTime)

	if err != nil {
		fmt.Println("Error while inserting ticket into ticket table")
		log.Fatal(err.Error())
	}

	return row
}
