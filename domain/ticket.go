package domain

import "time"

type TicketRequest struct {
	Name   string
	Phone  string
	PaxNum int
}

type TicketResponse struct {
	QueueId          string
	QueuePositionNum int
}

type Ticket struct {
	Id             string
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
