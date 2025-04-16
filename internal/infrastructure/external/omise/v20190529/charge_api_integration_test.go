//go:build integration
// +build integration

package v20190529_test

import (
	"go-tamboon/internal/infrastructure/external/omise"
	v20190529 "go-tamboon/internal/infrastructure/external/omise/v20190529"
	"log"
	"os"
	"testing"
)

func TestChargeApiIntegration(t *testing.T) {
	testCardToken := os.Getenv("OMISE_TEST_CARD_TOKEN")
	secretKey := os.Getenv("OMISE_API_SECRET_KEY")
	publicKey := os.Getenv("OMISE_API_PUBLIC_KEY")
	if secretKey == "" || publicKey == "" || testCardToken == "" {
		t.Skip("OMISE_API_SECRET_KEY, OMISE_API_PUBLIC_KEY or OMISE_TEST_CARD_TOKEN not set, test skipped")
	}
	config := omise.ApiConfig{
		SecretKey: secretKey,
		PublicKey: publicKey,
		Endpoints: omise.ApiEndpoints{
			Vault: "https://vault.omise.co/",
			Api:   "https://api.omise.co/",
		},
	}

	api := v20190529.NewChargeAPI(config)

	t.Run("Charge", func(t *testing.T) {
		t.Run("Should work", func(t *testing.T) {
			input := v20190529.ChargeRequest{
				Amount:   15000,
				Currency: "THB",
				Card:     testCardToken,
			}
			result, err := api.CreateCharge(input)

			log.Print(result)
			log.Print(err)
		})
	})
}
