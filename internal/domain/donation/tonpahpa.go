package donation

import "go-tamboon/shared/money"

type TonPahPa struct {
	total        money.Money
	faultyAmount money.Money
	donorCount   int
}

type TonPahPaSummary struct {
	Total      money.Money
	Average    money.Money
	Faulty     money.Money
	Successful money.Money
	DonorCount int
}

func NewTonPahPa(state TonPahPa) *TonPahPa {
	return &TonPahPa{
		total:        state.total,
		faultyAmount: state.faultyAmount,
		donorCount:   state.donorCount,
	}
}

func CreateTonPahPa() *TonPahPa {
	return &TonPahPa{
		total:        money.CreateNil(PahPaCurrency),
		faultyAmount: money.CreateNil(PahPaCurrency),
		donorCount:   0,
	}
}

func (t *TonPahPa) AddSongPahPa(songPahPa *SongPahPa) error {
	newTotal, err := t.total.Add(songPahPa.donateAmount)
	if err != nil {
		return err
	}
	t.total = newTotal

	if songPahPa.IsDonated() {
		t.donorCount++

		return nil
	}

	newFaultyAmount, err := t.faultyAmount.Add(songPahPa.donateAmount)
	if err != nil {
		return err
	}
	t.faultyAmount = newFaultyAmount

	return nil
}

// Summary
// Silencing the error from subtract and divided by are intended
// as no different currency and average when dividend is not zero respectively
func (t *TonPahPa) Summary() TonPahPaSummary {
	successful, _ := t.total.Subtract(t.faultyAmount)
	average := money.CreateNil(PahPaCurrency)
	if t.donorCount > 0 {
		newAverage, _ := successful.DividedBy(float64(t.donorCount))
		average = newAverage
	}

	return TonPahPaSummary{
		Total:      t.total,
		Successful: successful,
		Faulty:     t.faultyAmount,
		Average:    average,
		DonorCount: t.donorCount,
	}
}
