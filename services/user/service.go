package user

import "github.com/mateeullahmalik/goa-demo/internal/storage"

const (
	logPrefix = "user"
)

// Service represents loan service.
type Service struct {
	db storage.KeyValue
}

// NewService returns a new Service instance.
func NewService(db storage.KeyValue) *Service {
	return &Service{
		db: db,
	}
}
