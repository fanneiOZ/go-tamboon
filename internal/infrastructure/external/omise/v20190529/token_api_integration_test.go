//go:build integration
// +build integration

package v20190529_test

import (
	"fmt"
	"go-tamboon/internal/infrastructure/external/omise"
	v20190529 "go-tamboon/internal/infrastructure/external/omise/v20190529"
	"os"
	"testing"
)

func TestTokenApiIntegration(t *testing.T) {
	secretKey := os.Getenv("OMISE_API_SECRET_KEY")
	publicKey := os.Getenv("OMISE_API_PUBLIC_KEY")
	if secretKey == "" || publicKey == "" {
		t.Skip("OMISE_API_SECRET_KEY or OMISE_API_PUBLIC_KEY not set, test skipped")
	}
	config := omise.ApiConfig{
		SecretKey: secretKey,
		PublicKey: publicKey,
		Endpoints: omise.ApiEndpoints{
			Vault: "https://vault.omise.co/",
			Api:   "https://api.omise.co/",
		},
	}
	api := v20190529.NewTokenAPI(config)

	t.Run("CreateToken", func(t *testing.T) {
		t.Run("Should work", func(t *testing.T) {
			input := v20190529.TokenRequest{
				Object: "card",
				Card: v20190529.CardRequest{
					Name:            "Somchai Prasert",
					Number:          "4242424242424242",
					ExpirationMonth: 10,
					ExpirationYear:  2025,
				},
			}
			result, err := api.CreateToken(input)

			fmt.Println(result)
			fmt.Println(err)
		})
	})
}
