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
		{"GOLD_HAS_TOPTAN code should be HH_T", GOLD_HAS_TOPTAN, "HH_T"},
		{"GOLD_KULCE_TOPTAN code should be CH_T", GOLD_KULCE_TOPTAN, "CH_T"},
		{"GOLD_GRAM_TOPTAN code should be GAT", GOLD_GRAM_TOPTAN, "GAT"},
		{"GOLD_22_AYAR_HURDA code should be B_T", GOLD_22_AYAR_HURDA, "B_T"},
		{"GOLD_ATA_TOPTAN code should be A_T", GOLD_ATA_TOPTAN, "A_T"},
		{"GOLD_ESKI_CEYREK code should be EC", GOLD_ESKI_CEYREK, "EC"},
		{"GOLD_ESKI_YARIM code should be EY", GOLD_ESKI_YARIM, "EY"},
		{"GOLD_ESKI_TEKLIK code should be ET", GOLD_ESKI_TEKLIK, "ET"},
		{"GOLD_ESKI_GREMSE code should be EG", GOLD_ESKI_GREMSE, "EG"},
		{"GOLD_GUMUS code should be AG_T", GOLD_GUMUS, "AG_T"},
		{"GOLD_ATA_CUMHURIYET code should be A", GOLD_ATA_CUMHURIYET, "A"},
		{"GOLD_22_AYAR_BILEZIK code should be B", GOLD_22_AYAR_BILEZIK, "B"},
		{"GOLD_18_AYAR code should be 18", GOLD_18_AYAR, "18"},
		{"GOLD_14_AYAR code should be 14", GOLD_14_AYAR, "14"},
		{"GOLD_CEYREK code should be C", GOLD_CEYREK, "C"},
		{"GOLD_YARIM code should be Y", GOLD_YARIM, "Y"},
		{"GOLD_TEKLIK code should be T", GOLD_TEKLIK, "T"},
		{"GOLD_GREMSE code should be G", GOLD_GREMSE, "G"},
		{"GOLD_ATA_BESLI code should be A5", GOLD_ATA_BESLI, "A5"},
		{"GOLD_RESAT code should be R", GOLD_RESAT, "R"},
		{"GOLD_HAMIT code should be H", GOLD_HAMIT, "H"},
		{"GOLD_GRAM code should be GA", GOLD_GRAM, "GA"},
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
	c := a.GoldService.Get(GOLD_GUMUS)
	fmt.Println(c.Code)
	fmt.Println(c.Name)
	// Output:
	//AG_T
	//Gümüş
}
