package altinkaynak

import (
	"fmt"
	"testing"
)

func TestCurrencyGet(t *testing.T) {
	s := NewAltinkaynak().CurrencyService
	_ = s.Fetch()

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"CURRENCY_USD code should be USD", CurrencyUsd, "USD"},
		{"CURRENCY_EUR code should be EUR", CurrencyEur, "EUR"},
		{"CURRENCY_CHF code should be CHF", CurrencyChf, "CHF"},
		{"CURRENCY_GBP code should be GBP", CurrencyGbp, "GBP"},
		{"CURRENCY_JPY code should be JPY", CurrencyJpy, "JPY"},
		{"CURRENCY_SAR code should be SAR", CurrencySar, "SAR"},
		{"CURRENCY_AUD code should be AUD", CurrencyAud, "AUD"},
		{"CURRENCY_CAD code should be CAD", CurrencyCad, "CAD"},
		{"CURRENCY_RUB code should be RUB", CurrencyRub, "RUB"},
		{"CURRENCY_AZN code should be AZN", CurrencyAzn, "AZN"},
		{"CURRENCY_CNY code should be CNY", CurrencyCny, "CNY"},
		{"CURRENCY_RON code should be RON", CurrencyRon, "RON"},
		{"CURRENCY_AED code should be AED", CurrencyAed, "AED"},
		{"CURRENCY_KWD code should be KWD", CurrencyKwd, "KWD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.Get(tt.input).Code
			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func ExampleCurrencyService_Fetch() {
	a := NewAltinkaynak()
	_ = a.CurrencyService.Fetch()
}

func ExampleCurrencyService_Get() {
	a := NewAltinkaynak()
	_ = a.CurrencyService.Fetch()
	c := a.CurrencyService.Get(CurrencyUsd)
	fmt.Println(c.Code)
	fmt.Println(c.Name)
	// Output:
	//USD
	//Amerikan Doları
}
