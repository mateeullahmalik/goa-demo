package loan

import (
	"time"

	"github.com/mateeullahmalik/goa-demo/domain/loan"
	"github.com/mateeullahmalik/goa-demo/internal/storage"
)

const (
	logPrefix = "loan"
)

// Service represents loan service.
type Service struct {
	db storage.KeyValue
}

func (service *Service) GetLoan(loanID int) *loan.Loan {
	// get from db

	// returning dummy obj
	currTime := time.Now()
	return &loan.Loan{
		ID:          1,
		Amount:      100.0,
		PaybackDate: &currTime,
		LenderID:    2,
		BorrowerID:  3,
	}
}

// NewService returns a new Service instance.
func NewService(db storage.KeyValue) *Service {
	return &Service{
		db: db,
	}
}
