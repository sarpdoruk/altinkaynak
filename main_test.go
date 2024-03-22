package altinkaynak

import (
	"fmt"
	"testing"
)

func TestMainGet(t *testing.T) {
	s := NewAltinkaynak().MainService
	_ = s.Fetch()

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"CURRENCY_USD code should be USD", CURRENCY_USD, "USD"},
		{"CURRENCY_EUR code should be EUR", CURRENCY_EUR, "EUR"},
		{"GOLD_HAS_TOPTAN code should be HH_T", GOLD_HAS_TOPTAN, "HH_T"},
		{"GOLD_KULCE_TOPTAN code should be CH_T", GOLD_KULCE_TOPTAN, "CH_T"},
		{"GOLD_22_AYAR_HURDA code should be B_T", GOLD_22_AYAR_HURDA, "B_T"},
		{"GOLD_GUMUS code should be AG_T", GOLD_GUMUS, "AG_T"},
		{"GOLD_CEYREK code should be C", GOLD_CEYREK, "C"},
		{"GOLD_ONS code should be ONS", GOLD_ONS, "ONS"},
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

func ExampleMainService_Fetch() {
	a := NewAltinkaynak()
	_ = a.MainService.Fetch()
}

func ExampleMainService_Get() {
	a := NewAltinkaynak()
	_ = a.MainService.Fetch()
	c := a.MainService.Get(GOLD_CEYREK)
	fmt.Println(c.Code)
	fmt.Println(c.Name)
	// Output:
	//C
	//Çeyrek Altın
}
