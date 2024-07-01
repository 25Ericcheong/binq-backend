package main

import (
	"database/sql"
	"errors"
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
	Id             string
	Branch         string
	CustomerName   string
	CustomerPaxNum int
	CustomerPhone  string
}

type DbTicket struct {
	Id                string
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

	newTicket := Ticket{"1", "Damansara", "Eric", 5, "0122817216"}
	newTicket1 := Ticket{"1", "Damansara", "Bobby", 1, "0122817216"}
	newTicket2 := Ticket{"1", "Damansara", "Billy", 3, "0122817216"}

	row, exists := getTicket(db, newTicket.Id)

	if !exists {
		row = insertDbTicket(db, newTicket)
	}

	insertDbTicket(db, newTicket1)
	insertDbTicket(db, newTicket2)

	fmt.Printf("Ticket details \n"+
		"Id: %s \n"+
		"Customer Name: %s \n"+
		"Creation Date: %s \n", row.Id, row.CustomerName, row.CreatedOnDateTime)

	tickets, err := getTicketsByBranch(db, "Damansara")
	if err != nil {
		log.Println("Error while trying to read multiple tickets from a branch")
		log.Fatal(err.Error())
	}

	fmt.Println("Going through inserted tickets found in ticket based on branch: Damansara")
	for _, ticket := range tickets {
		fmt.Println(ticket.Id + " " + ticket.CustomerName + " " + ticket.Branch)
	}

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

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println("Error occurred while trying to close database")
			log.Fatal(err.Error())
		}
	}(db)
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
	query := `CREATE TABLE IF NOT EXISTS ticket
		(id SERIAL PRIMARY KEY, 
		branch VARCHAR(50) NOT NULL,
		customer_name VARCHAR(100) NOT NULL,
		customer_pax_num INTEGER NOT NULL CHECK(customer_pax_num > 0),
		customer_phone VARCHAR(20) NOT NULL,
		created_on_date_time TIMESTAMP NOT NULL DEFAULT NOW())`

	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error while creating ticket table")
		log.Fatal(err.Error())
	}
}

func insertDbTicket(db *sql.DB, inputTicket Ticket) (ticket DbTicket) {
	query := `INSERT INTO ticket 
    	(branch, customer_name, customer_pax_num, customer_phone) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, branch, customer_name, customer_pax_num, customer_phone, created_on_date_time`

	err := db.
		QueryRow(query, inputTicket.Branch, inputTicket.CustomerName, inputTicket.CustomerPaxNum, inputTicket.CustomerPhone).
		Scan(&ticket.Id, &ticket.Branch, &ticket.CustomerName, &ticket.CustomerPaxNum, &ticket.CustomerPhone, &ticket.CreatedOnDateTime)

	if err != nil {
		fmt.Println("Error while inserting ticket into ticket table")
		log.Fatal(err.Error())
	}

	return ticket
}

func getTicket(db *sql.DB, ticketId string) (ticket DbTicket, exists bool) {
	query := `SELECT * FROM ticket WHERE id = $1`

	err := db.
		QueryRow(query, ticketId).
		Scan(&ticket.Id, &ticket.Branch, &ticket.CustomerName, &ticket.CustomerPaxNum, &ticket.CustomerPhone, &ticket.CreatedOnDateTime)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No row found with provided id: " + ticketId)
			return DbTicket{}, false
		}

		//Probably no rows found
		fmt.Println("Unexpected error " + err.Error())
		return DbTicket{}, false
	}

	return ticket, true
}

func deleteTicket(db *sql.DB, ticketId string) error {
	query := `DELETE FROM ticket WHERE id = $1`

	var err = db.QueryRow(query, ticketId).Scan()
	if err != nil {
		return err
	}

	return nil
}

func getTicketsByBranch(db *sql.DB, branch string) ([]DbTicket, error) {
	query := `SELECT * FROM ticket WHERE branch = $1`

	rows, err := db.Query(query, branch)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error while trying to close rows of tickets acquired from database")
		}
	}(rows)

	var tickets []DbTicket

	for rows.Next() {
		var ticket DbTicket
		err := rows.Scan(&ticket.Id, &ticket.Branch, &ticket.CustomerName, &ticket.CustomerPhone,
			&ticket.CustomerPaxNum, &ticket.CreatedOnDateTime)

		if err != nil {
			return tickets, err
		}
		tickets = append(tickets, ticket)
	}

	if rows.Err() != nil {
		return tickets, err
	}
	return tickets, nil
}
