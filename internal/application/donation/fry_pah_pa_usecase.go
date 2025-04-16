package donation

import (
	"encoding/csv"
	"errors"
	"fmt"
	paymentApp "go-tamboon/internal/application/payment"
	"go-tamboon/internal/domain/donation"
	"go-tamboon/internal/domain/payment"
	"go-tamboon/shared/money"
	"io"
	"strconv"
	"strings"
)

var (
	expectedHeaders = [6]string{
		"Name",
		"AmountSubunits",
		"CCNumber",
		"CVV",
		"ExpMonth",
		"ExpYear",
	}
)

type FryPahPaUseCase struct {
	tonPahPa      *donation.TonPahPa
	chargeUseCase paymentApp.ChargeCreditCard
}

func NewFryPahPaUseCase(chargeCreditCardUseCase paymentApp.ChargeCreditCard) *FryPahPaUseCase {
	return &FryPahPaUseCase{
		tonPahPa:      donation.CreateTonPahPa(),
		chargeUseCase: chargeCreditCardUseCase,
	}
}

func (useCase *FryPahPaUseCase) Execute(inputReader io.Reader) (donation.TonPahPaSummary, error) {
	csvReader := csv.NewReader(inputReader)
	headerRow, err := csvReader.Read()
	if err != nil {
		return donation.TonPahPaSummary{}, err
	}

	if !validateHeaderRow(headerRow) {
		return donation.TonPahPaSummary{}, fmt.Errorf("unexpected csv headers: %w", err)
	}

	rowNumber := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		rowNumber++

		if err != nil {
			// log error and skip
			continue
		}

		pahPaDto, err := parseRecord(record)
		if err != nil {
			// parse data error: log and skip
			continue
		}

		if err := useCase.donate(pahPaDto, rowNumber); err != nil {
			// donation error: log and skip
			continue
		}
	}

	return useCase.tonPahPa.Summary(), nil
}

func (useCase *FryPahPaUseCase) donate(input PahPaDto, rowNumber int) error {
	amount, err := strconv.ParseFloat(input.AmountSubunits, 64)
	if err != nil {
		return fmt.Errorf("parse donation amount error: %w", err)
	}

	inputExpMonth, err := strconv.ParseUint(input.ExpMonth, 10, 8)
	if err != nil {
		return fmt.Errorf("parse credit card expMonth error: %w", err)
	}

	inputExpYear, err := strconv.ParseInt(input.ExpYear, 10, 16)
	if err != nil {
		return fmt.Errorf("parse credit card expYear error: %w", err)
	}

	donateByCard, err := payment.CreateCard(
		input.CCNumber,
		input.Name,
		input.CVV,
		uint8(inputExpMonth),
		int(inputExpYear),
	)
	if err != nil {
		return fmt.Errorf("donating card is invalid: %w", err)
	}

	songPahPa, err := donation.CreateSongPahPa(
		strconv.FormatInt(int64(rowNumber), 10),
		input.Name,
		money.New(donation.PahPaCurrency, amount/100),
		donateByCard,
	)
	if err != nil {
		return fmt.Errorf("unable to create song-pah-pa: %w", err)
	}

	transaction := useCase.chargeUseCase.Execute(
		songPahPa.DonateByCard(),
		songPahPa.DonateAmount().Amount(),
		songPahPa.DonateAmount().Currency(),
	)
	songPahPa.AttachTransaction(transaction)
	err = useCase.tonPahPa.AddSongPahPa(songPahPa)
	if err != nil {
		return fmt.Errorf("unable to attach song-pah-pa: %w", err)
	}

	return nil
}

func validateHeaderRow(record []string) bool {
	if len(record) != len(expectedHeaders) {
		return false
	}

	for i := 0; i < len(expectedHeaders); i++ {
		if record[i] != expectedHeaders[i] {
			return false
		}
	}

	return true
}

func parseRecord(record []string) (PahPaDto, error) {
	if len(record) < 6 {
		return PahPaDto{}, errors.New("record is too short")
	}

	return PahPaDto{
		Name:           strings.TrimSpace(record[0]),
		AmountSubunits: strings.TrimSpace(record[1]),
		CCNumber:       strings.TrimSpace(record[2]),
		CVV:            strings.TrimSpace(record[3]),
		ExpMonth:       strings.TrimSpace(record[4]),
		ExpYear:        strings.TrimSpace(record[5]),
	}, nil
}
