package v20190529

type ChargeRequest struct {
	Amount   int    //`json:"amount"`
	Currency string //`json:"currency"`
	Card     string //`json:"card"`
}

type TokenRequest struct {
	Object string      //`json:"object"` // always "card"
	Card   CardRequest //`json:"card"`
}

type CardRequest struct {
	Name            string  `json:"name"`
	Number          string  `json:"number"`
	ExpirationMonth int     `json:"expiration_month"`
	ExpirationYear  int     `json:"expiration_year"`
	SecurityCode    *string `json:"security_code,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	Email           *string `json:"email,omitempty"`
	PhoneNumber     *string `json:"phone_number,omitempty"`
	PostalCode      *string `json:"postal_code,omitempty"`
	State           *string `json:"state,omitempty"`
	Street1         *string `json:"street1,omitempty"`
	Street2         *string `json:"street2,omitempty"`
}
