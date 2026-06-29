package altinkaynak

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// goldFixture is a deterministic snapshot of the gold endpoint response so the
// tests don't depend on the live (and occasionally flaky) Altinkaynak API.
const goldFixture = `[{"Alis":"6.313,18","Satis":"6.360,31","Kod":"HH_T","Aciklama":"Has Toptan","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"6.281,62","Satis":"6.331,04","Kod":"CH_T","Aciklama":"Külçe Toptan","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"6.281,62","Satis":"6.401,29","Kod":"GAT","Aciklama":"Gram Toptan","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"41.855,47","Satis":"44.600,00","Kod":"PA","Aciklama":"Ata Cumhuriyet","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"5.757,43","Satis":"5.800,99","Kod":"B_T","Aciklama":"22 Ayar Hurda","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"41.855,47","Satis":"42.735,65","Kod":"A_T","Aciklama":"Ata Toptan","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"5.757,43","Satis":"6.210,00","Kod":"PB","Aciklama":"22 Ayar Bilezik","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"6.281,62","Satis":"6.395,00","Kod":"PGA","Aciklama":"Gram Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"10.238,04","Satis":"10.850,00","Kod":"PC","Aciklama":"Çeyrek Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"20.311,93","Satis":"21.700,00","Kod":"PY","Aciklama":"Yarım Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"40.729,24","Satis":"43.400,00","Kod":"PT","Aciklama":"Teklik Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"101.286,06","Satis":"108.500,00","Kod":"PG","Aciklama":"Gremse Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"41.396,30","Satis":"44.600,00","Kod":"PR","Aciklama":"Resat Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"41.396,30","Satis":"44.600,00","Kod":"PH","Aciklama":"Hamit Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"206.397,12","Satis":"223.000,00","Kod":"PA5","Aciklama":"Ata Beşli","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"3.576,51","Satis":"4.050,00","Kod":"P14","Aciklama":"14 Ayar Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"4.417,65","Satis":"5.050,00","Kod":"P18","Aciklama":"18 Ayar Altın","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"10.244,12","Satis":"10.484,94","Kod":"EC","Aciklama":"Eski Çeyrek","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"20.307,88","Satis":"20.804,51","Kod":"EY","Aciklama":"Eski Yarım","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"40.692,76","Satis":"41.743,77","Kod":"ET","Aciklama":"Eski Teklik","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"101.245,53","Satis":"104.328,79","Kod":"EG","Aciklama":"Eski Gremse","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"41.571,62","Satis":"43.139,37","Kod":"EA","Aciklama":"Eski Ata","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"95,45","Satis":"107,65","Kod":"AG_T","Aciklama":"Gümüş","GuncellenmeZamani":"19.06.2026 16:10:44"},{"Alis":"4.181,50","Satis":"4.181,50","Kod":"XAUUSD","Aciklama":"ONS","GuncellenmeZamani":"19.06.2026 16:10:00"},{"Alis":"6.385,00","Satis":"6.385,00","Kod":"IAB_KAPANIS","Aciklama":"İAB Kapanış","GuncellenmeZamani":"18.06.2026 00:00:00"}]`

// newFixtureServer returns a test server that serves the given JSON body with a
// JSON content-type, mimicking the Altinkaynak endpoints.
func newFixtureServer(body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(body))
	}))
}

func TestGoldGet(t *testing.T) {
	srv := newFixtureServer(goldFixture)
	defer srv.Close()

	s := NewAltinkaynak(WithApiUrl("", srv.URL)).GoldService
	if err := s.Fetch(); err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"GoldHasToptan code should be HH_T", GoldHasToptan, "HH_T"},
		{"GoldKulceToptan code should be CH_T", GoldKulceToptan, "CH_T"},
		{"GoldAtaCumhuriyet code should be PA", GoldAtaCumhuriyet, "PA"},
		{"GoldGramToptan code should be GAT", GoldGramToptan, "GAT"},
		{"Gold22AyarHurda code should be B_T", Gold22AyarHurda, "B_T"},
		{"GoldAtaToptan code should be A_T", GoldAtaToptan, "A_T"},
		{"Gold22AyarBilezik code should be PB", Gold22AyarBilezik, "PB"},
		{"Gold18Ayar code should be P18", Gold18Ayar, "P18"},
		{"Gold14Ayar code should be P14", Gold14Ayar, "P14"},
		{"GoldCeyrek code should be PC", GoldCeyrek, "PC"},
		{"GoldYarim code should be PY", GoldYarim, "PY"},
		{"GoldTeklik code should be PT", GoldTeklik, "PT"},
		{"GoldGremse code should be PG", GoldGremse, "PG"},
		{"GoldAtaBesli code should be PA5", GoldAtaBesli, "PA5"},
		{"GoldResat code should be PR", GoldResat, "PR"},
		{"GoldHamit code should be PH", GoldHamit, "PH"},
		{"GoldGram code should be PGA", GoldGram, "PGA"},
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
	// In real usage the default constructor points at the live API:
	//   a := NewAltinkaynak()
	//   _ = a.GoldService.Fetch()
	a := NewAltinkaynak()
	_ = a.GoldService.Fetch()
	c := a.GoldService.Get(GoldGumus)
	fmt.Println(c.Code)
}
