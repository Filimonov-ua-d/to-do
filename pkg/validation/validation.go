package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"interview.go/models"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

// ValidateCustomer validates the customer model.
func (v Validator) ValidateCustomer(customer *models.Customer) error {
	return validation.ValidateStruct(customer,
		validation.Field(&customer.FirstName, validation.Required, validation.Length(1, 50)),
		validation.Field(&customer.LastName, validation.Required, validation.Length(1, 50)),
		validation.Field(&customer.Email, validation.Required, is.Email),
		validation.Field(&customer.Gender),
		validation.Field(&customer.IpAddress, validation.Required, is.IPv4))
}
