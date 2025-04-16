package payment

import (
	v20190529 "go-tamboon/internal/infrastructure/external/omise/v20190529"
	"math"
)

var _ PaymentService = (*OmisePaymentService)(nil)

type OmisePaymentService struct {
	tokenApi  *v20190529.TokenAPI
	chargeApi *v20190529.ChargeAPI
}

func NewOmisePaymentService(tokenApi *v20190529.TokenAPI, chargeApi *v20190529.ChargeAPI) *OmisePaymentService {
	return &OmisePaymentService{tokenApi: tokenApi, chargeApi: chargeApi}
}

func (s *OmisePaymentService) Gateway() PaymentGateway {
	return GatewayOmise
}

func (s *OmisePaymentService) Charge(card Card, transaction *Transaction) error {
	if transaction == nil {
		return ErrTransactionNotFound
	}

	tokenResponse, err := s.tokenApi.CreateToken(v20190529.TokenRequest{
		Object: "card",
		Card: v20190529.CardRequest{
			Number:          card.Number(),
			Name:            card.Holder(),
			ExpirationMonth: int(card.expiryMonth),
			ExpirationYear:  int(card.expiryYear),
		},
	})
	if err != nil {
		return transaction.MarkFailed(err.Error())
	}

	cardToken := tokenResponse.Id
	chargeResponse, err := s.chargeApi.CreateCharge(v20190529.ChargeRequest{
		Amount:   int(math.Round(transaction.paymentAmount.Amount() * 100)),
		Currency: transaction.paymentAmount.Currency(),
		Card:     cardToken,
	})
	if err != nil {
		return transaction.MarkFailed(err.Error())
	}

	return transaction.MarkSucceeded(chargeResponse.Id)
}
