package money

import (
	"errors"
	"github.com/shopspring/decimal"
)

type Money struct {
	currency string
	amount   decimal.Decimal
}

func New(currency string, amount float64) Money {
	return Money{currency: currency, amount: decimal.NewFromFloat(amount).Round(2)}
}

func CreateNil(currency string) Money {
	return Money{currency: currency, amount: decimal.Zero}
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Amount() float64 {
	amount, _ := m.amount.Float64()

	return amount
}

func (m Money) String() string {
	return m.amount.StringFixed(2)
}

var (
	ErrMismatchedCurrency = errors.New("money currencies do not match")
	ErrDivideByZero       = errors.New("money cannot be divided by zero")
)

func (m Money) Add(addend Money) (Money, error) {
	if m.currency != addend.currency {
		return Money{}, ErrMismatchedCurrency
	}

	return Money{m.currency, m.amount.Add(addend.amount)}, nil
}

func (m Money) Subtract(operand Money) (Money, error) {
	if operand.currency != m.currency {
		return Money{}, ErrMismatchedCurrency
	}

	return Money{m.currency, m.amount.Sub(operand.amount)}, nil
}

func (m Money) MultipliedBy(multiplier float64) Money {
	return Money{m.currency, m.amount.Mul(decimal.NewFromFloat(multiplier)).Round(2)}
}

func (m Money) DividedBy(dividend float64) (Money, error) {
	if dividend == 0 {
		return Money{}, ErrDivideByZero
	}

	return Money{m.currency, m.amount.Div(decimal.NewFromFloat(dividend)).Round(2)}, nil
}
