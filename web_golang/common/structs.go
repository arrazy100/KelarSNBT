package common

import "fmt"

type ErrorValidationResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (repo ErrorValidationResponse) String() string {
	return fmt.Sprintf("{ failed_field: %s, tag: %s, value: %s }", repo.FailedField, repo.Tag, repo.Value)
}
