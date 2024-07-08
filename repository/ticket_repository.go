package repository

import (
	"context"
	"database/sql"
	"github.com/25Ericcheong/binq-backend/domain"
	"time"
)

type TicketRepository interface {
	CreateTicket(ctx context.Context, newTicket domain.Ticket) (TicketDb, error)
	GetTicketById(ctx context.Context, ticketId string) (TicketDb, error)
	GetTicketsByBranch(ctx context.Context, branch string) ([]TicketDb, error)
	UpdateTicket(ctx context.Context, ticketId string, updatedTicket domain.Ticket) error
	DeleteTicket(ctx context.Context, ticketId string) error
}

type TicketDb struct {
	Id                string
	Branch            string
	CustomerName      string
	CustomerPaxNum    int
	CustomerPhone     string
	CreatedOnDateTime time.Time
}

type psqlTicketRepository struct {
	DB *sql.DB
}

func NewTicketRepository(DB *sql.DB) TicketRepository {
	return &psqlTicketRepository{
		DB: DB,
	}
}

func (t *psqlTicketRepository) CreateTicket(ctx context.Context, newTicket domain.Ticket) (TicketDb, error) {
	query := `INSERT INTO ticket 
    	(branch, customer_name, customer_pax_num, customer_phone) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, branch, customer_name, customer_pax_num, customer_phone, created_on_date_time`

	createdTicket := TicketDb{}
	err := t.DB.
		QueryRow(query, newTicket.Branch, newTicket.CustomerName, newTicket.CustomerPaxNum, newTicket.CustomerPhone).
		Scan(&createdTicket.Id, &createdTicket.Branch, &createdTicket.CustomerName, &createdTicket.CustomerPaxNum, &createdTicket.CustomerPhone, &createdTicket.CreatedOnDateTime)

	if err != nil {
		return TicketDb{}, err
	}

	return createdTicket, nil
}

func (t *psqlTicketRepository) GetTicketById(ctx context.Context, ticketId string) (TicketDb, error) {
	query := `SELECT * FROM ticket WHERE id = $1`

	ticket := TicketDb{}
	err := t.DB.
		QueryRow(query, ticketId).
		Scan(&ticket.Id, &ticket.Branch, &ticket.CustomerName, &ticket.CustomerPaxNum, &ticket.CustomerPhone, &ticket.CreatedOnDateTime)

	if err != nil {
		return TicketDb{}, nil
	}

	return ticket, nil
}

func (t *psqlTicketRepository) GetTicketsByBranch(ctx context.Context, branch string) ([]TicketDb, error) {
	query := `SELECT * FROM ticket WHERE branch = $1`

	rows, err := t.DB.Query(query, branch)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
	}(rows)

	var tickets []TicketDb
	for rows.Next() {
		var ticket TicketDb
		err := rows.Scan(
			&ticket.Id,
			&ticket.Branch,
			&ticket.CustomerName,
			&ticket.CustomerPhone,
			&ticket.CustomerPaxNum,
			&ticket.CreatedOnDateTime,
		)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return tickets, nil
}

func (t *psqlTicketRepository) UpdateTicket(ctx context.Context, ticketId string, updatedTicket domain.Ticket) error {
	query := `UPDATE ticket SET branch = $2, customer_name = $3, customer_pax_num = $4, customer_phone = $5
              WHERE id = $1`

	var err = t.DB.QueryRow(query, ticketId,
		updatedTicket.Branch, updatedTicket.CustomerName, updatedTicket.CustomerPaxNum, updatedTicket.CustomerPhone).Scan()

	if err != nil {
		return err
	}

	return nil
}

func (t *psqlTicketRepository) DeleteTicket(ctx context.Context, ticketId string) error {
	query := `DELETE FROM ticket WHERE id = $1`

	var err = t.DB.QueryRow(query, ticketId).Scan()
	if err != nil {
		return err
	}

	return nil
}
