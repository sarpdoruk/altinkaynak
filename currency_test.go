package altinkaynak

import (
	"fmt"
	"testing"
)

// currencyFixture is a deterministic snapshot of the currency endpoint response
// so the tests don't depend on the live Altinkaynak API.
const currencyFixture = `[{"Alis":"46,490","Satis":"46,650","Kod":"USD","Aciklama":"Amerikan Doları","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"52,893","Satis":"53,180","Kod":"EUR","Aciklama":"Avrupa Para Birimi","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"56,901","Satis":"57,650","Kod":"CHF","Aciklama":"İsviçre Frangı","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"61,221","Satis":"61,893","Kod":"GBP","Aciklama":"İngiliz Sterlini","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"0,2832","Satis":"0,2889","Kod":"JPY","Aciklama":"Japon Yeni","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"12,218","Satis":"12,429","Kod":"SAR","Aciklama":"S. Arabistan Riyali","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"31,428","Satis":"32,156","Kod":"AUD","Aciklama":"Avustralya Doları","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"32,469","Satis":"32,958","Kod":"CAD","Aciklama":"Kanada Doları","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"0,5382","Satis":"0,6421","Kod":"RUB","Aciklama":"Rus Rublesi","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"25,940","Satis":"27,533","Kod":"AZN","Aciklama":"Azerbaycan Manatı","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"6,233","Satis":"7,272","Kod":"CNY","Aciklama":"Çin Yuanı","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"8,989","Satis":"10,278","Kod":"RON","Aciklama":"Romanya Leyi","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"11,946","Satis":"13,343","Kod":"AED","Aciklama":"B.A.E. Dirhemi","GuncellenmeZamani":"29.06.2026 13:47:54"},{"Alis":"139,695","Satis":"153,551","Kod":"KWD","Aciklama":"Kuveyt Dinarı","GuncellenmeZamani":"29.06.2026 13:47:54"}]`

func TestCurrencyGet(t *testing.T) {
	srv := newFixtureServer(currencyFixture)
	defer srv.Close()

	s := NewAltinkaynak(WithApiUrl(srv.URL, "")).CurrencyService
	if err := s.Fetch(); err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

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
}
