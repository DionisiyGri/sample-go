package postgres

import (
	"database/sql"
)

// SomeService implement SomeServiceManager interface for data manipulations
type SomeService struct {
	db *sql.DB
}

// NewSomeService creates new instance of SomeService for psql db
func NewSomeService(db *sql.DB) *SomeService {
	return &SomeService{db: db}
}
