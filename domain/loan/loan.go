package loan

import (
	"time"
)

// Loan ...
type Loan struct {
	ID          int
	Amount      float64
	LenderID    int
	BorrowerID  int
	Description *string
	PaybackDate *time.Time
}
