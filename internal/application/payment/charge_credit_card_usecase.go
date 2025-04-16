package payment

import (
	"go-tamboon/internal/domain/payment"
	"go-tamboon/shared/money"
)

type ChargeCreditCard struct {
	paymentService payment.PaymentService
}

func NewChargeCreditCard(paymentService payment.PaymentService) *ChargeCreditCard {
	return &ChargeCreditCard{paymentService: paymentService}
}

func (useCase *ChargeCreditCard) Execute(inputCard payment.Card, amount float64, currency string) *payment.Transaction {
	chargingAmount := money.New(currency, amount)
	transaction := payment.CreateTransaction(chargingAmount)
	chargingCard, err := payment.NewCard(inputCard)
	if err != nil {
		_ = transaction.MarkFailed(err.Error())

		return transaction
	}

	useCase.paymentService.Charge(chargingCard, transaction)

	return transaction
}
