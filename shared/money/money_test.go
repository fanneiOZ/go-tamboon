package money_test

import (
	"errors"
	"go-tamboon/shared/money"
	"reflect"
	"testing"
)

func TestMoney(t *testing.T) {
	t.Parallel()
	t.Run("New", func(t *testing.T) {
		t.Run("Should return new Money with 2-decimal places amount and currency", func(t *testing.T) {
			result := money.New("thb", 234.1259)

			if result.Amount() != 234.13 {
				t.Errorf("Money.Amount() should be 234.13, got %v", result.Amount())
			}
			if result.String() != "234.13" {
				t.Errorf("Money.String() should be \"234.13\", got %s", result.String())
			}

			if result.Currency() != "thb" {
				t.Errorf("Money.Currency() should be thb, got %s", result.Currency())
			}
		})

		t.Run("Should have trailing decimals value when extract the String", func(t *testing.T) {
			result := money.New("thb", 234)

			if result.String() != "234.00" {
				t.Errorf("Money.String() should be \"234.00\", got %s", result.String())
			}
		})

		t.Run("Should maintain equality value objects", func(t *testing.T) {
			object1 := money.New("thb", 234.1259)
			object2 := money.New("thb", 234.1259)

			if !reflect.DeepEqual(object1, object2) {
				t.Errorf("Money.New() should be %v, got %v", object1, object2)
			}

			if object1 == object2 {
				t.Errorf("Money.New() should not be equal directly")
			}
		})
	})

	t.Run("CreateNil", func(t *testing.T) {
		t.Run("Should return new Money with 2-decimal places amount and currency", func(t *testing.T) {
			result := money.CreateNil("thb")

			if result.Amount() != 0 {
				t.Errorf("Money.Amount() should be 0, got %v", result.Amount())
			}
			if result.String() != "0.00" {
				t.Errorf("Money.String() should be \"0\", got %s", result.String())
			}

			if result.Currency() != "thb" {
				t.Errorf("Money.Currency() should be thb, got %s", result.Currency())
			}
		})
	})

	t.Run("Add", func(t *testing.T) {
		t.Run("Should add the amount correctly", func(t *testing.T) {
			a := money.New("thb", 234.1259)
			b := money.New("thb", 234.1259)
			expected := money.New("thb", 468.26)
			result, _ := a.Add(b)

			if result.Currency() != expected.Currency() {
				t.Errorf("result.Currency() should be %v, got %v", expected.Currency(), result.Currency())
			}
			if result.Amount() != expected.Amount() {
				t.Errorf("result.Amount() should be %v, got %v", expected.Amount(), result.Amount())
			}
		})

		t.Run("Should return error when currency mismatched", func(t *testing.T) {
			a := money.New("thb", 234.1259)
			b := money.New("vnd", 234.1259)
			expectedError := errors.New("money currencies do not match")
			result, actualError := a.Add(b)

			if actualError.Error() != expectedError.Error() {
				t.Errorf("Money.Add() should return error %v, got %v", expectedError, actualError)
			}

			if result.Amount() != 0 {
				t.Errorf("Fallback result.Amount() should be 0, got %v", result.Amount())
			}

			if result.Currency() != "" {
				t.Errorf("Fallback result.Currency() should be thb, got %v", result.Currency())
			}
		})
	})

	t.Run("Subtract", func(t *testing.T) {
		t.Run("Should subtract the amount correctly", func(t *testing.T) {
			a := money.New("thb", 7595.1259)
			b := money.New("thb", 234.1259)
			expected := money.New("thb", 7361)
			result, _ := a.Subtract(b)

			if result.Currency() != expected.Currency() {
				t.Errorf("result.Currency() should be %v, got %v", expected.Currency(), result.Currency())
			}
			if result.Amount() != expected.Amount() {
				t.Errorf("result.Amount() should be %v, got %v", expected.Amount(), result.Amount())
			}
		})

		t.Run("Should subtract the amount correctly when operand more than original", func(t *testing.T) {
			a := money.New("thb", 234.1259)
			b := money.New("thb", 7595.1259)
			expected := money.New("thb", -7361)
			result, _ := a.Subtract(b)

			if result.Currency() != expected.Currency() {
				t.Errorf("result.Currency() should be %v, got %v", expected.Currency(), result.Currency())
			}
			if result.Amount() != expected.Amount() {
				t.Errorf("result.Amount() should be %v, got %v", expected.Amount(), result.Amount())
			}
		})

		t.Run("Should return error when currency mismatched", func(t *testing.T) {
			a := money.New("thb", 234.1259)
			b := money.New("vnd", 234.1259)
			expectedError := errors.New("money currencies do not match")
			result, actualError := a.Subtract(b)

			if actualError.Error() != expectedError.Error() {
				t.Errorf("Money.Subtract() should return error %v, got %v", expectedError, actualError)
			}

			if result.Amount() != 0 {
				t.Errorf("Fallback result.Amount() should be 0, got %v", result.Amount())
			}

			if result.Currency() != "" {
				t.Errorf("Fallback result.Currency() should be thb, got %v", result.Currency())
			}
		})
	})

	t.Run("MultipliedBy", func(t *testing.T) {
		t.Run("Should multiply the amount correctly", func(t *testing.T) {
			a := money.New("thb", 234.1259)
			expected := money.New("thb", 11735.3)
			result := a.MultipliedBy(50.123)

			if result.Currency() != expected.Currency() {
				t.Errorf("result.Currency() should be %v, got %v", expected.Currency(), result.Currency())
			}
			if result.Amount() != expected.Amount() {
				t.Errorf("result.Amount() should be %v, got %v", expected.Amount(), result.Amount())
			}
		})
	})

	t.Run("DividedBy", func(t *testing.T) {
		t.Run("Should divide the amount correctly", func(t *testing.T) {
			a := money.New("thb", 234.1259)
			expected := money.New("thb", 4.67)
			result, _ := a.DividedBy(50.123)

			if result.Currency() != expected.Currency() {
				t.Errorf("result.Currency() should be %v, got %v", expected.Currency(), result.Currency())
			}
			if result.Amount() != expected.Amount() {
				t.Errorf("result.Amount() should be %v, got %v", expected.Amount(), result.Amount())
			}
		})

		t.Run("Should return error when divided by zero", func(t *testing.T) {
			a := money.New("thb", 234.1259)
			expectedError := errors.New("money cannot be divided by zero")
			result, actualError := a.DividedBy(0)

			if actualError.Error() != expectedError.Error() {
				t.Errorf("Money.DividedBy() should return error %v, got %v", expectedError, actualError)
			}

			if result.Amount() != 0 {
				t.Errorf("Fallback result.Amount() should be 0, got %v", result.Amount())
			}

			if result.Currency() != "" {
				t.Errorf("Fallback result.Currency() should be thb, got %v", result.Currency())
			}
		})
	})
}
