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
	var ticket domain.CreateTicketResponse
	ticketInput := domain.Ticket{
		Branch:         newTicket.Branch,
		CustomerName:   newTicket.CustomerName,
		CustomerPaxNum: newTicket.CustomerPaxNum,
		CustomerPhone:  newTicket.CustomerPhone,
	}

	ticketDb, err := t.ticketRepo.CreateTicket(ctx, ticketInput)
	if err != nil {
		return ticket, err
	}

	ticket = domain.CreateTicketResponse{
		Id:             ticketDb.Id,
		Branch:         ticketDb.Branch,
		CustomerName:   ticketDb.CustomerName,
		CustomerPhone:  ticketDb.CustomerPhone,
		CustomerPaxNum: ticketDb.CustomerPaxNum,
	}

	return ticket, nil
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

func (t *ticketUseCase) UpdateTicket(ctx context.Context, updatedTicket domain.UpdateTicketRequest) error {
	ticketInput := domain.Ticket{
		Branch:         updatedTicket.Branch,
		CustomerName:   updatedTicket.CustomerName,
		CustomerPhone:  updatedTicket.CustomerPhone,
		CustomerPaxNum: updatedTicket.CustomerPaxNum,
	}

	err := t.ticketRepo.UpdateTicket(ctx, updatedTicket.Id, ticketInput)
	if err != nil {
		return err
	}

	return nil
}

func (t *ticketUseCase) DeleteTicket(ctx context.Context, ticketId string) error {
	err := t.ticketRepo.DeleteTicket(ctx, ticketId)
	if err != nil {
		return err
	}

	return nil
}
