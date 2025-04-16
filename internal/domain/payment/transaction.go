package payment

import (
	"errors"
	"go-tamboon/shared/money"
	"strings"
	"time"
)

var (
	ErrMissingExternalRefId = errors.New("missing external ref_id")
	ErrMissingFailureDetail = errors.New("missing failure detail")
)

type TransactionStatus string

const (
	TransactionStatusCreated   TransactionStatus = "created"
	TransactionStatusSucceeded TransactionStatus = "succeeded"
	TransactionStatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
	status         TransactionStatus
	paymentAmount  money.Money
	externalRefId  string
	failureMessage string
	createdAt      time.Time
	updatedAt      time.Time
}

func NewTransaction(state Transaction) *Transaction {
	return &Transaction{
		status:         state.status,
		paymentAmount:  state.paymentAmount,
		externalRefId:  state.externalRefId,
		failureMessage: state.failureMessage,
		createdAt:      state.createdAt,
		updatedAt:      state.updatedAt,
	}
}

func CreateTransaction(amount money.Money) *Transaction {
	return &Transaction{
		status:        TransactionStatusCreated,
		paymentAmount: amount,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}
}

func (t *Transaction) PaymentAmount() money.Money {
	return t.paymentAmount
}

func (t *Transaction) Status() TransactionStatus {
	return t.status
}

func (t *Transaction) IsPaymentSuccessful() bool {
	return t.status == TransactionStatusSucceeded
}

func (t *Transaction) MarkFailed(failureMessage string) error {
	if strings.TrimSpace(failureMessage) == "" {
		return ErrMissingFailureDetail
	}

	t.status = TransactionStatusFailed
	t.failureMessage = failureMessage
	t.updatedAt = time.Now()

	return nil
}

func (t *Transaction) MarkSucceeded(externalRefId string) error {
	if strings.TrimSpace(externalRefId) == "" {
		return ErrMissingExternalRefId
	}

	t.status = TransactionStatusSucceeded
	t.externalRefId = externalRefId
	t.updatedAt = time.Now()

	return nil
}
