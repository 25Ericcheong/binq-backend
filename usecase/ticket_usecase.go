package usecase

import (
	"context"
	"github.com/25Ericcheong/binq-backend/domain"
)

type ticketUseCase struct {
	ticketRepo domain.TicketRepository
}

func NewTicketUseCase(t domain.TicketRepository) domain.TicketUseCase {
	return &ticketUseCase{
		ticketRepo: t,
	}
}

func (t *ticketUseCase) CreateTicket(ctx context.Context, newTicket domain.CreateTicketRequest) (domain.TicketResponse, error) {

	ticket := domain.Ticket{
		Branch:         newTicket.Branch,
		CustomerName:   newTicket.CustomerName,
		CustomerPaxNum: newTicket.CustomerPaxNum,
		CustomerPhone:  newTicket.CustomerPhone,
	}

	ticketDb, err := t.ticketRepo.CreateTicket(ctx, ticket)
	if err != nil {
		return domain.TicketResponse{}, err
	}

	return domain.TicketResponse{
		Id:                ticketDb.Id,
		CustomerName:      ticketDb.CustomerName,
		CustomerPhone:     ticketDb.CustomerPhone,
		CustomerPaxNum:    ticketDb.CustomerPaxNum,
		CreatedOnDateTime: ticketDb.CreatedOnDateTime,
	}, nil
}

func (t *ticketUseCase) GetTicketByBranch(ctx context.Context, branch string) ([]domain.TicketResponse, error) {
	var ticketsRes []domain.TicketResponse

	ticketDbsBranchSpecific, err := t.ticketRepo.GetTicketsByBranch(ctx, branch)
	if err != nil {
		return ticketsRes, err
	}

	for _, ticketDb := range ticketDbsBranchSpecific {
		ticketRes := domain.TicketResponse{
			Id:                ticketDb.Id,
			CustomerName:      ticketDb.CustomerName,
			CustomerPhone:     ticketDb.CustomerPhone,
			CustomerPaxNum:    ticketDb.CustomerPaxNum,
			CreatedOnDateTime: ticketDb.CreatedOnDateTime,
		}
		ticketsRes = append(ticketsRes, ticketRes)
	}

	return ticketsRes, nil
}
