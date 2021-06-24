// Code generated by goa v3.4.3, DO NOT EDIT.
//
// loan client
//
// Command:
// $ goa gen github.com/mateeullahmalik/goa-demo/api/design

package loan

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "loan" service client.
type Client struct {
	ListLoansEndpoint goa.Endpoint
	GetLoanEndpoint   goa.Endpoint
}

// NewClient initializes a "loan" service client given the endpoints.
func NewClient(listLoans, getLoan goa.Endpoint) *Client {
	return &Client{
		ListLoansEndpoint: listLoans,
		GetLoanEndpoint:   getLoan,
	}
}

// ListLoans calls the "listLoans" endpoint of the "loan" service.
func (c *Client) ListLoans(ctx context.Context) (res LoanCollection, err error) {
	var ires interface{}
	ires, err = c.ListLoansEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(LoanCollection), nil
}

// GetLoan calls the "getLoan" endpoint of the "loan" service.
// GetLoan may return the following errors:
//	- "not_found" (type *NotFound): Loan not found
//	- error: internal error
func (c *Client) GetLoan(ctx context.Context, p *GetLoanPayload) (res *Loan, err error) {
	var ires interface{}
	ires, err = c.GetLoanEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Loan), nil
}
