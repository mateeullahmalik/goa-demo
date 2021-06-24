package services

import (
	"context"
	"fmt"

	"github.com/mateeullahmalik/goa-demo/api"
	"github.com/mateeullahmalik/goa-demo/internal/log"
	"github.com/mateeullahmalik/goa-demo/internal/storage"
	"github.com/mateeullahmalik/goa-demo/internal/storage/memory"

	"github.com/mateeullahmalik/goa-demo/api/gen/http/loan/server"
	"github.com/mateeullahmalik/goa-demo/api/gen/loan"
	loansrvc "github.com/mateeullahmalik/goa-demo/services/loan"

	goahttp "goa.design/goa/v3/http"
)

// LoanSrvc represents services for user endpoints.
type LoanSrvc struct {
	*Common
	loan *loansrvc.Service
	db   storage.KeyValue
	// other assets
}

// List all stored loans
func (service *LoanSrvc) ListLoans(context.Context) (res loan.LoanCollection, err error) {
	fmt.Println("ListLoan called")

	return
}

// get loan by ID
// The "view" return value must have one of the following views
//	- "default"
//	- "tiny"
func (service *LoanSrvc) GetLoan(_ context.Context, p *loan.GetLoanPayload) (res *loan.Loan, view string, err error) {
	l := service.loan.GetLoan(p.ID)

	paybackDate := l.PaybackDate.String()
	res = &loan.Loan{
		ID:          l.ID,
		Amount:      l.Amount,
		LenderID:    l.LenderID,
		BorrowerID:  l.BorrowerID,
		PaybackDate: &paybackDate,
	}

	return res, p.View, nil
}

// Mount configures the mux to serve the user endpoints.
func (service *LoanSrvc) Mount(ctx context.Context, mux goahttp.Muxer) goahttp.Server {
	endpoints := loan.NewEndpoints(service)
	srv := server.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, api.ErrorHandler, nil)
	server.Mount(mux, srv)

	for _, m := range srv.Mounts {
		log.WithContext(ctx).Infof("%q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	return srv
}

// NewLoanSrvc returns the loan service implementation.
func NewLoanSrvc(srvc *loansrvc.Service) *LoanSrvc {
	return &LoanSrvc{
		Common: NewCommon(),
		loan:   srvc,
		db:     memory.NewKeyValue(),
	}
}
