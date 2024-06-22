package api

type TicketRequest struct {
	Name   string
	Phone  string
	PaxNum int
}

type TicketResponse struct {
	QueueId          string
	QueuePositionNum int
}
