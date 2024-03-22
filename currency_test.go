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
		{"CURRENCY_USD code should be USD", CURRENCY_USD, "USD"},
		{"CURRENCY_EUR code should be EUR", CURRENCY_EUR, "EUR"},
		{"CURRENCY_CHF code should be CHF", CURRENCY_CHF, "CHF"},
		{"CURRENCY_GBP code should be GBP", CURRENCY_GBP, "GBP"},
		{"CURRENCY_JPY code should be JPY", CURRENCY_JPY, "JPY"},
		{"CURRENCY_SAR code should be SAR", CURRENCY_SAR, "SAR"},
		{"CURRENCY_AUD code should be AUD", CURRENCY_AUD, "AUD"},
		{"CURRENCY_CAD code should be CAD", CURRENCY_CAD, "CAD"},
		{"CURRENCY_RUB code should be RUB", CURRENCY_RUB, "RUB"},
		{"CURRENCY_AZN code should be AZN", CURRENCY_AZN, "AZN"},
		{"CURRENCY_CNY code should be CNY", CURRENCY_CNY, "CNY"},
		{"CURRENCY_RON code should be RON", CURRENCY_RON, "RON"},
		{"CURRENCY_AED code should be AED", CURRENCY_AED, "AED"},
		{"CURRENCY_BGN code should be BGN", CURRENCY_BGN, "BGN"},
		{"CURRENCY_KWD code should be KWD", CURRENCY_KWD, "KWD"},
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
	c := a.CurrencyService.Get(CURRENCY_USD)
	fmt.Println(c.Code)
	fmt.Println(c.Name)
	// Output:
	//USD
	//Amerikan Doları
}
