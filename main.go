package main

import (
	"go-tamboon/internal/application/donation"
	paymentApp "go-tamboon/internal/application/payment"
	"go-tamboon/internal/domain/payment"
	"go-tamboon/internal/infrastructure/external/omise"
	v20190529 "go-tamboon/internal/infrastructure/external/omise/v20190529"
	"go-tamboon/internal/infrastructure/fileio"
	"go-tamboon/internal/presenter/console"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <path to .rot128 file>", os.Args[0])
	}

	path := os.Args[1]
	decodedReader, err := fileio.OpenAndDecodeRot128File(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := decodedReader.Close(); err != nil {
			log.Printf("Warning: failed to close file: %s", err)
		}
	}()

	fryPahPaUseCase := initFryPahPaUseCase()
	summary, err := fryPahPaUseCase.Execute(decodedReader.Reader)
	if err != nil {
		log.Fatal(err)
	}

	console.PrintTonPahPaSummary(summary)
}

func initFryPahPaUseCase() *donation.FryPahPaUseCase {
	omiseApiConfig := omise.ApiConfig{
		SecretKey: os.Getenv("OMISE_API_SECRET_KEY"),
		PublicKey: os.Getenv("OMISE_API_PUBLIC_KEY"),
		Endpoints: omise.ApiEndpoints{
			Api:   "https://api.omise.co/",
			Vault: "https://vault.omise.co/",
		},
	}

	omiseTokenApi := v20190529.NewTokenAPI(omiseApiConfig)
	omiseChargeApi := v20190529.NewChargeAPI(omiseApiConfig)
	omisePaymentService := payment.NewOmisePaymentService(omiseTokenApi, omiseChargeApi)
	chargeCreditCardUseCase := paymentApp.NewChargeCreditCard(omisePaymentService)

	return donation.NewFryPahPaUseCase(chargeCreditCardUseCase)
}
