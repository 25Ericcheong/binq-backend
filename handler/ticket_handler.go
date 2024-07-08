package handler

import (
	"context"
	"github.com/25Ericcheong/binq-backend/usecase"
)

type TicketHandlers interface {
	CreateTicket(ctx context.Context) error
	GetTicketsByBranch(ctx context.Context) error
	UpdateTicket(ctx context.Context) error
	DeleteTicket(ctx context.Context) error
}

type ticketHandlers struct {
	t usecase.TicketUseCase
}

func NewHttpHandler(ticketUseCase usecase.TicketUseCase) TicketHandlers {
	return &ticketHandlers{t: ticketUseCase}
}

func (t *ticketHandlers) CreateTicket(ctx context.Context) error {
	return nil
}

func (t *ticketHandlers) GetTicketsByBranch(ctx context.Context) error {
	return nil
}

func (t *ticketHandlers) UpdateTicket(ctx context.Context) error {
	return nil
}

func (t *ticketHandlers) DeleteTicket(ctx context.Context) error {
	return nil
}
