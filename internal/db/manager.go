package db

import "github.com/DionisiyGri/sample-go/internal/model"

// SomeServiceManager provide an interface for data manipulations
type SomeServiceManager interface {
	GetAll() ([]model.Some, error)
	PostSome(*model.Some) (*model.Some, error)
}
