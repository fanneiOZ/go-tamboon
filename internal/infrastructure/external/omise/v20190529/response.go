package v20190529

type BaseApiObject struct {
	Object    string `json:"object"`
	Id        string `json:"id"`
	LiveMode  bool   `json:"livemode"`
	Location  string `json:"location"`
	CreatedAt string `json:"created_at"`
}

type CardResponse struct {
	BaseApiObject
	Deleted            bool    `json:"deleted"`
	Street1            string  `json:"street1"`
	Street2            string  `json:"street2"`
	City               string  `json:"city"`
	State              string  `json:"state"`
	PhoneNumber        string  `json:"phone_number"`
	PostalCode         string  `json:"postal_code"`
	Country            string  `json:"country"`
	Financing          string  `json:"financing"`
	Bank               string  `json:"bank"`
	Brand              string  `json:"brand"`
	Fingerprint        string  `json:"fingerprint"`
	FirstDigits        string  `json:"first_digits"`
	LastDigits         string  `json:"last_digits"`
	Name               string  `json:"name"`
	ExpirationMonth    int     `json:"expiration_month"`
	ExpirationYear     int     `json:"expiration_year"`
	SecurityCodeCheck  bool    `json:"security_code_check"`
	TokenizationMethod *string `json:"tokenization_method"` // nullable string
}

type TokenResponse struct {
	BaseApiObject
	Used         bool         `json:"used"`
	ChargeStatus string       `json:"charge_status"` // enum + unknown
	Card         CardResponse `json:"card"`
}

type ChargeStatus string

const (
	StatusFailed     ChargeStatus = "failed"
	StatusExpired    ChargeStatus = "expired"
	StatusPending    ChargeStatus = "pending"
	StatusReversed   ChargeStatus = "reversed"
	StatusSuccessful ChargeStatus = "successful"
	StatusUnknown    ChargeStatus = "unknown"
)

type ChargeResponse struct {
	BaseApiObject
	Amount          int `json:"amount"`
	Net             int `json:"net"`
	Fee             int `json:"fee"`
	FeeVAT          int `json:"fee_vat"`
	Interest        int `json:"interest"`
	InterestVAT     int `json:"interest_vat"`
	FundingAmount   int `json:"funding_amount"`
	RefundedAmount  int `json:"refunded_amount"`
	TransactionFees struct {
		FeeFlat string `json:"fee_flat"`
		FeeRate string `json:"fee_rate"`
		VATRate string `json:"vat_rate"`
	} `json:"transaction_fees"`
	PlatformFee struct {
		Fixed      int `json:"fixed"`
		Amount     int `json:"amount"`
		Percentage int `json:"percentage"`
	} `json:"platform_fee"`
	Currency        string         `json:"currency"`
	FundingCurrency string         `json:"funding_currency"`
	IP              string         `json:"ip"`
	Metadata        map[string]any `json:"metadata"`
	Description     string         `json:"description"`
	Card            CardResponse   `json:"card"`

	Transaction    string       `json:"transaction"`
	FailureCode    string       `json:"failure_code"`
	FailureMessage string       `json:"failure_message"`
	Status         ChargeStatus `json:"status"`

	PaidAt     string `json:"paid_at"`
	ExpiresAt  string `json:"expires_at"`
	ExpiredAt  string `json:"expired_at"`
	ReversedAt string `json:"reversed_at"`
	Authorized bool   `json:"authorized"`
	Capturable bool   `json:"capturable"`
	Refundable bool   `json:"refundable"`
	Reversible bool   `json:"reversible"`
	Voided     bool   `json:"voided"`
	Paid       bool   `json:"paid"`
	Expired    bool   `json:"expired"`
}
