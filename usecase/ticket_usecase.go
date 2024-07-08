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

func (t *ticketUseCase) CreateTicket(ctx context.Context, newTicket domain.CreateTicketRequest) (domain.CreateTicketResponse, error) {

	ticket := domain.Ticket{
		Branch:         newTicket.Branch,
		CustomerName:   newTicket.CustomerName,
		CustomerPaxNum: newTicket.CustomerPaxNum,
		CustomerPhone:  newTicket.CustomerPhone,
	}

	ticketDb, err := t.ticketRepo.CreateTicket(ctx, ticket)
	if err != nil {
		return domain.CreateTicketResponse{}, err
	}

	return domain.CreateTicketResponse{
		Id:             ticketDb.Id,
		Branch:         ticketDb.Branch,
		CustomerName:   ticketDb.CustomerName,
		CustomerPhone:  ticketDb.CustomerPhone,
		CustomerPaxNum: ticketDb.CustomerPaxNum,
	}, nil
}

func (t *ticketUseCase) GetTicketsByBranch(ctx context.Context, branch string) ([]domain.GetTicketByBranchResponse, error) {
	var ticketsByBranch []domain.GetTicketByBranchResponse

	ticketDbsByBranch, err := t.ticketRepo.GetTicketsByBranch(ctx, branch)
	if err != nil {
		return ticketsByBranch, err
	}

	for _, ticketDb := range ticketDbsByBranch {
		ticketRes := domain.GetTicketByBranchResponse{
			Id:                ticketDb.Id,
			CustomerName:      ticketDb.CustomerName,
			CustomerPhone:     ticketDb.CustomerPhone,
			CustomerPaxNum:    ticketDb.CustomerPaxNum,
			CreatedOnDateTime: ticketDb.CreatedOnDateTime,
		}
		ticketsByBranch = append(ticketsByBranch, ticketRes)
	}

	return ticketsByBranch, nil
}

func (t *ticketUseCase) UpdateTicket(ctx context.Context, ticketId string) error {
	return nil
}

func (t *ticketUseCase) DeleteTicket(ctx context.Context, ticketId string) error {
	return nil
}
