package repos

import (
	"github.com/DionisiyGri/sample-go/internal/db"
	"github.com/DionisiyGri/sample-go/internal/db/postgres"
)

// SomeRepository - handler for some entities for different dbs (if you will add others)
type SomeRepository struct {
	Postgres db.SomeServiceManager
}

// NewSomeRepository builds new SomeRepository object
func NewSomeRepository(postgres *postgres.SomeService) *SomeRepository {
	return &SomeRepository{
		Postgres: postgres,
	}
}
