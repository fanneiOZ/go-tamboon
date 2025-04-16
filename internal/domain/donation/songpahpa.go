package donation

import (
	"go-tamboon/internal/domain/payment"
	"go-tamboon/shared/money"
	"strings"
)

type SongPahPa struct {
	id           string
	donorName    string
	donateAmount money.Money
	transaction  *payment.Transaction
	donateByCard payment.Card
}

func NewSongPahPa(state SongPahPa) *SongPahPa {
	return &SongPahPa{
		id:           state.id,
		donorName:    state.donorName,
		donateAmount: state.donateAmount,
		transaction:  state.transaction,
		donateByCard: state.donateByCard,
	}
}

func CreateSongPahPa(id string, donorName string, amount money.Money, donateByCard payment.Card) (*SongPahPa, error) {
	inputId := strings.TrimSpace(id)
	inputDonorName := strings.TrimSpace(donorName)
	if inputId == "" || inputDonorName == "" {
		return nil, ErrRequiredFieldMissing
	}

	if amount.Amount() <= 0 {
		return nil, ErrInvalidDonateAmount
	}

	return &SongPahPa{
		id:           inputId,
		donorName:    inputDonorName,
		donateAmount: amount,
		transaction:  &payment.Transaction{},
		donateByCard: donateByCard,
	}, nil
}

func (s *SongPahPa) Id() string {
	return s.id
}

func (s *SongPahPa) DonorName() string {
	return s.donorName
}

func (s *SongPahPa) DonateAmount() money.Money {
	return s.donateAmount
}

func (s *SongPahPa) DonateByCard() payment.Card {
	return s.donateByCard
}

func (s *SongPahPa) IsDonated() bool {
	return s.transaction.IsPaymentSuccessful()
}

func (s *SongPahPa) AttachTransaction(transaction *payment.Transaction) *SongPahPa {
	s.transaction = transaction

	return s
}
