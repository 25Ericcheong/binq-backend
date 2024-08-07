package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/25Ericcheong/binq-backend/domain"
	"github.com/25Ericcheong/binq-backend/repository"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	// to be able to open postgres driver
	_ "github.com/lib/pq"
)

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createTicketTable(db)

	ticketRepository := repository.NewTicketRepository(db)

	//newTicket := domain.Ticket{Branch: "Damansara", CustomerName: "Eric", CustomerPaxNum: 5, CustomerPhone: "0122817216"}
	newTicket1 := domain.Ticket{Branch: "Damansara", CustomerName: "Bobby", CustomerPaxNum: 1, CustomerPhone: "0122817216"}
	newTicket2 := domain.Ticket{Branch: "Damansara", CustomerName: "Billy", CustomerPaxNum: 3, CustomerPhone: "0122817216"}

	//row, err := ticketRepository.GetTicketById(ctx, newTicket.Id)
	//if err != nil {
	//	log.Println("Error while trying to get ticket " + newTicket.Id)
	//	log.Fatal(err.Error())
	//}

	_, err = ticketRepository.CreateTicket(ctx, newTicket1)
	if err != nil {
		log.Println("Error while trying to create a ticket")
		log.Fatal(err.Error())
	}

	row, err := ticketRepository.CreateTicket(ctx, newTicket2)
	if err != nil {
		log.Println("Error while trying to create a ticket")
		log.Fatal(err.Error())
	}
	fmt.Printf("Ticket details \n"+
		"Id: %s \n"+
		"Customer Name: %s \n"+
		"Creation Date: %s \n", row.Id, row.CustomerName, row.CreatedOnDateTime)

	tickets, err := ticketRepository.GetTicketsByBranch(ctx, "Damansara")
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
