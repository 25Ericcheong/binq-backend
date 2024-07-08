package domain

import (
	"context"
	"time"
)

type TicketRepository interface {
	CreateTicket(ctx context.Context, newTicket Ticket) (TicketDb, error)
	GetTicketById(ctx context.Context, ticketId string) (TicketDb, error)
	GetTicketsByBranch(ctx context.Context, branch string) ([]TicketDb, error)
	UpdateTicket(ctx context.Context, ticketId string) error
	DeleteTicket(ctx context.Context, ticketId string) error
}

type TicketUseCase interface {
	CreateTicket(ctx context.Context, newTicket CreateTicketRequest) (TicketResponse, error)
	GetTicketsByBranch(ctx context.Context, branch string) ([]TicketResponse, error)
	UpdateTicket(ctx context.Context, ticketId string) error
	DeleteTicket(ctx context.Context, ticketId string) error
}

type CreateTicketRequest struct {
	Branch         string
	CustomerName   string
	CustomerPaxNum int
	CustomerPhone  string
}

type TicketResponse struct {
	Id                string
	Branch            string
	CustomerName      string
	CustomerPaxNum    int
	CustomerPhone     string
	CreatedOnDateTime time.Time
}

type Ticket struct {
	Branch         string
	CustomerName   string
	CustomerPaxNum int
	CustomerPhone  string
}

type TicketDb struct {
	Id                string
	Branch            string
	CustomerName      string
	CustomerPaxNum    int
	CustomerPhone     string
	CreatedOnDateTime time.Time
}