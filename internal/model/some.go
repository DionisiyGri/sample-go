package model

import (
	"github.com/DionisiyGri/sample-go/internal/validation"
)

type Some struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty" valid:"length(1|30)"`
}

// Validate implements interface Validateble from common/validator
func (s *Some) Validate() (bool, validation.Errors) {
	errors := validation.Errors{}

	if s.Name == "" {
		errors["name"] = messageFieldShouldNotBeEmpty
	}

	if len(errors) != 0 {
		return false, errors
	}

	return true, errors
}
