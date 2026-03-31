package altinkaynak

import (
	"fmt"
	"testing"
)

func TestGoldGet(t *testing.T) {
	s := NewAltinkaynak().GoldService
	_ = s.Fetch()

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"GoldHasToptan code should be HH_T", GoldHasToptan, "HH_T"},
		{"GoldKulceToptan code should be CH_T", GoldKulceToptan, "CH_T"},
		{"GoldAtaCumhuriyet code should be A", GoldAtaCumhuriyet, "A"},
		{"GoldGramToptan code should be GAT", GoldGramToptan, "GAT"},
		{"Gold22AyarHurda code should be B_T", Gold22AyarHurda, "B_T"},
		{"GoldAtaToptan code should be A_T", GoldAtaToptan, "A_T"},
		{"Gold22AyarBilezik code should be B", Gold22AyarBilezik, "B"},
		{"Gold18Ayar code should be 18", Gold18Ayar, "18"},
		{"Gold14Ayar code should be 14", Gold14Ayar, "14"},
		{"GoldCeyrek code should be C", GoldCeyrek, "C"},
		{"GoldYarim code should be Y", GoldYarim, "Y"},
		{"GoldTeklik code should be T", GoldTeklik, "T"},
		{"GoldGremse code should be G", GoldGremse, "G"},
		{"GoldAtaBesli code should be A5", GoldAtaBesli, "A5"},
		{"GoldResat code should be R", GoldResat, "R"},
		{"GoldHamit code should be H", GoldHamit, "H"},
		{"GoldGram code should be GA", GoldGram, "GA"},
		{"GoldEskiCeyrek code should be EC", GoldEskiCeyrek, "EC"},
		{"GoldEskiYarim code should be EY", GoldEskiYarim, "EY"},
		{"GoldEskiTeklik code should be ET", GoldEskiTeklik, "ET"},
		{"GoldEskiGremse code should be EG", GoldEskiGremse, "EG"},
		{"GoldGumus code should be AG_T", GoldGumus, "AG_T"},
		{"GoldOns code should be XAUUSD", GoldOns, "XAUUSD"},
		{"GoldIabKapanis code should be IAB_KAPANIS", GoldIabKapanis, "IAB_KAPANIS"},
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

func ExampleGoldService_Fetch() {
	a := NewAltinkaynak()
	_ = a.GoldService.Fetch()
}

func ExampleGoldService_Get() {
	a := NewAltinkaynak()
	_ = a.GoldService.Fetch()
	c := a.GoldService.Get(GoldGumus)
	fmt.Println(c.Code)
	fmt.Println(c.Name)
	// Output:
	//AG_T
	//Gümüş
}
