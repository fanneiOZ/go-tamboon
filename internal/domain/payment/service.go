package payment

type PaymentService interface {
	Gateway() PaymentGateway
	Charge(card Card, transaction *Transaction) error
}
