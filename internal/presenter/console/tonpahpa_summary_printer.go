package console

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"go-tamboon/internal/domain/donation"
	"go-tamboon/shared/money"
	"strings"
)

func PrintTonPahPaSummary(summary donation.TonPahPaSummary) {
	printLine := func(label string, inputAmount money.Money) {
		amount := humanize.FormatFloat("#,###.##", inputAmount.Amount())
		fmt.Printf("%20s: %-5s %15s\n", label, strings.ToUpper(inputAmount.Currency()), amount)
	}

	printLine("total received", summary.Total)
	printLine("successfully donated", summary.Successful)
	printLine("faulty donation", summary.Faulty)
	fmt.Println()
	printLine("average per person", summary.Average)
}
