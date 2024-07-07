package repository

import (
	"context"
	"database/sql"
	"github.com/25Ericcheong/binq-backend/domain"
)

type psqlTicketRepository struct {
	DB *sql.DB
}

func NewTicketRepository(DB *sql.DB) domain.TicketRepository {
	return &psqlTicketRepository{
		DB: DB,
	}
}

func (t *psqlTicketRepository) CreateTicket(ctx context.Context, newTicket domain.Ticket) (domain.TicketDb, error) {
	query := `INSERT INTO ticket 
    	(branch, customer_name, customer_pax_num, customer_phone) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, branch, customer_name, customer_pax_num, customer_phone, created_on_date_time`

	createdTicket := domain.TicketDb{}
	err := t.DB.
		QueryRow(query, newTicket.Branch, newTicket.CustomerName, newTicket.CustomerPaxNum, newTicket.CustomerPhone).
		Scan(&createdTicket.Id, &createdTicket.Branch, &createdTicket.CustomerName, &createdTicket.CustomerPaxNum, &createdTicket.CustomerPhone, &createdTicket.CreatedOnDateTime)

	if err != nil {
		return domain.TicketDb{}, err
	}

	return createdTicket, nil
}

func (t *psqlTicketRepository) GetTicketById(ctx context.Context, ticketId string) (domain.TicketDb, error) {
	query := `SELECT * FROM ticket WHERE id = $1`

	ticket := domain.TicketDb{}
	err := t.DB.
		QueryRow(query, ticketId).
		Scan(&ticket.Id, &ticket.Branch, &ticket.CustomerName, &ticket.CustomerPaxNum, &ticket.CustomerPhone, &ticket.CreatedOnDateTime)

	if err != nil {
		return domain.TicketDb{}, nil
	}

	return ticket, nil
}

func (t *psqlTicketRepository) GetTicketsByBranch(ctx context.Context, branch string) ([]domain.TicketDb, error) {
	query := `SELECT * FROM ticket WHERE branch = $1`

	rows, err := t.DB.Query(query, branch)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
	}(rows)

	var tickets []domain.TicketDb
	for rows.Next() {
		var ticket domain.TicketDb
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

func (t *psqlTicketRepository) UpdateTicket(ctx context.Context, ticketId string) error {
	query := `UPDATE ticket SET branch = $2, customer_name = $3, customer_pax_num = $4, customer_phone = $5
              WHERE id = $1`

	updatedTicket := domain.TicketDb{}
	var err = t.DB.QueryRow(query, updatedTicket.Id,
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
