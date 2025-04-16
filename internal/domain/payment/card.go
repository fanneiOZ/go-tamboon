package payment

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidExpiryMonth   = errors.New("invalid expiry month")
	ErrInvalidExpiryYear    = errors.New("invalid expiry year")
	ErrFieldValueIsRequired = errors.New("field value is required")
)

var bangkokTimezone, _ = time.LoadLocation("Asia/Bangkok")

type Card struct {
	number       string
	holder       string
	securityCode string
	expiryMonth  uint8
	expiryYear   int16
}

func CreateCard(number string, holder string, securityCode string, expiryMonth uint8, expiryYear int) (Card, error) {
	card, err := NewCard(Card{
		number:       number,
		holder:       holder,
		securityCode: securityCode,
		expiryMonth:  expiryMonth,
		expiryYear:   int16(expiryYear),
	})

	if err != nil {
		return Card{}, err
	}

	return card, nil
}

func NewCard(state Card) (Card, error) {
	if !(state.expiryMonth >= 1 && state.expiryMonth <= 12) {
		return Card{}, ErrInvalidExpiryMonth
	}

	if !(state.expiryYear >= 2000 && state.expiryYear <= 9999) {
		return Card{}, ErrInvalidExpiryYear
	}

	inputNumber := strings.TrimSpace(state.number)
	inputHolder := strings.TrimSpace(state.holder)
	inputSecurityCode := strings.TrimSpace(state.securityCode)
	if inputNumber == "" || inputHolder == "" || inputSecurityCode == "" {
		return Card{}, ErrFieldValueIsRequired
	}

	return Card{
		number:       inputNumber,
		holder:       inputHolder,
		securityCode: inputSecurityCode,
		expiryMonth:  state.expiryMonth,
		expiryYear:   state.expiryYear,
	}, nil
}

func (c Card) Number() string {
	return c.number
}

func (c Card) Holder() string {
	return c.holder
}

func (c Card) SecurityCode() string {
	return c.securityCode
}

func (c Card) ExpiryMonth() string {
	return fmt.Sprintf("%02d", c.expiryMonth)
}

func (c Card) ExpiryYear() string {
	return strconv.FormatInt(int64(c.expiryYear), 10)
}

func (c Card) IsExpired() bool {
	expiryToCheck := int(c.expiryYear)*1000 + int(c.expiryMonth)
	currentDate := time.Now().In(bangkokTimezone)
	currentFiscalMonth := currentDate.Year()*1000 + int(currentDate.Month())

	return currentFiscalMonth > expiryToCheck
}
