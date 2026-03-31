// Package altinkaynak provides services for fetching currency and gold data from [Altinkaynak API].
//
// [Altinkaynak API]: https://www.altinkaynak.com/Araclar/Servisler
package altinkaynak

import (
	"io"
	"net/http"
	"time"
)

const (
	dateTimeFormat = "02.1.2006 15:04:05"
)

// Currency constants
const (
	CurrencyUsd string = "USD"
	CurrencyEur        = "EUR"
	CurrencyChf        = "CHF"
	CurrencyGbp        = "GBP"
	CurrencyJpy        = "JPY"
	CurrencySar        = "SAR"
	CurrencyAud        = "AUD"
	CurrencyCad        = "CAD"
	CurrencyRub        = "RUB"
	CurrencyAzn        = "AZN"
	CurrencyCny        = "CNY"
	CurrencyRon        = "RON"
	CurrencyAed        = "AED"
	CurrencyKwd        = "KWD"
)

// Gold constants
const (
	GoldHasToptan     string = "HH_T"
	GoldKulceToptan          = "CH_T"
	GoldAtaCumhuriyet        = "A"
	GoldGramToptan           = "GAT"
	Gold22AyarHurda          = "B_T"
	GoldAtaToptan            = "A_T"
	Gold22AyarBilezik        = "B"
	Gold18Ayar               = "18"
	Gold14Ayar               = "14"
	GoldCeyrek               = "C"
	GoldYarim                = "Y"
	GoldTeklik               = "T"
	GoldGremse               = "G"
	GoldAtaBesli             = "A5"
	GoldResat                = "R"
	GoldHamit                = "H"
	GoldGram                 = "GA"
	GoldEskiCeyrek           = "EC"
	GoldEskiYarim            = "EY"
	GoldEskiTeklik           = "ET"
	GoldEskiGremse           = "EG"
	GoldGumus                = "AG_T"
	GoldOns                  = "XAUUSD"
	GoldIabKapanis           = "IAB_KAPANIS"
)

var (
	location *time.Location
)

// ServiceInterface represents the interface for the Altinkaynak API
type ServiceInterface interface {
	Fetch() error
	Get(code string) Resource
}

// config represents the configuration for the Altinkaynak API
type config struct {
	CurrencyServiceApiUrl string
	GoldServiceApiUrl     string
}

// Option represents a functional option for the Altinkaynak API
type Option func(*config)

// Altinkaynak represents the Altinkaynak API
type Altinkaynak struct {
	CurrencyService *CurrencyService
	GoldService     *GoldService
}

// Resource represents a currency or gold resource
type Resource struct {
	Code         string `json:"Kod"`
	Name         string `json:"Aciklama"`
	BuyString    string `json:"Alis"`
	SellString   string `json:"Satis"`
	Buy          float64
	Sell         float64
	UpdatedAtRaw string `json:"GuncellenmeZamani"`
	UpdatedAt    time.Time
}

// SendRequest sends a request to the given URL with the given payload
func SendRequest(method string, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	return body, nil
}

func WithApiUrl(CurrencyServiceApiUrl, GoldServiceApiUrl string) Option {
	return func(c *config) {
		c.CurrencyServiceApiUrl = CurrencyServiceApiUrl
		c.GoldServiceApiUrl = GoldServiceApiUrl
	}
}

func NewAltinkaynak(opts ...Option) *Altinkaynak {
	// Set location
	location, _ = time.LoadLocation("Europe/Istanbul")

	// Default credentials
	config := &config{
		CurrencyServiceApiUrl: "https://static.altinkaynak.com/public/Currency",
		GoldServiceApiUrl:     "https://static.altinkaynak.com/public/Gold",
	}

	// Apply options
	for _, opt := range opts {
		opt(config)
	}

	// Return Altinkaynak instance
	return &Altinkaynak{
		CurrencyService: &CurrencyService{
			apiUrl: config.CurrencyServiceApiUrl,
		},
		GoldService: &GoldService{
			apiUrl: config.GoldServiceApiUrl,
		},
	}
}
