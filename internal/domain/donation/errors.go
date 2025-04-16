package donation

import "errors"

var (
	ErrRequiredFieldMissing = errors.New("required field missing")
	ErrInvalidDonateAmount  = errors.New("donate amount must be greater than zero")
)
