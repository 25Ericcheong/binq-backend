package domain

import (
	"time"
)

type CreateTicketRequest struct {
	Branch         string
	CustomerName   string
	CustomerPaxNum int
	CustomerPhone  string
}

type CreateTicketResponse struct {
	Id             string
	Branch         string
	CustomerName   string
	CustomerPaxNum int
	CustomerPhone  string
}

type GetTicketByBranchResponse struct {
	Id                string
	Branch            string
	CustomerName      string
	CustomerPaxNum    int
	CustomerPhone     string
	CreatedOnDateTime time.Time
}

type UpdateTicketRequest struct {
	Id             string
	Branch         string
	CustomerName   string
	CustomerPaxNum int
	CustomerPhone  string
}

type Ticket struct {
	Branch         string
	CustomerName   string
	CustomerPaxNum int
	CustomerPhone  string
}
