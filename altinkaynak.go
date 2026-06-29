// Package altinkaynak provides services for fetching currency and gold data from [Altinkaynak API].
//
// [Altinkaynak API]: https://www.altinkaynak.com/Araclar/Servisler
package altinkaynak

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	GoldAtaCumhuriyet        = "PA"
	GoldGramToptan           = "GAT"
	Gold22AyarHurda          = "B_T"
	GoldAtaToptan            = "A_T"
	Gold22AyarBilezik        = "PB"
	Gold18Ayar               = "P18"
	Gold14Ayar               = "P14"
	GoldCeyrek               = "PC"
	GoldYarim                = "PY"
	GoldTeklik               = "PT"
	GoldGremse               = "PG"
	GoldAtaBesli             = "PA5"
	GoldResat                = "PR"
	GoldHamit                = "PH"
	GoldGram                 = "PGA"
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
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")
	// A browser-like User-Agent reduces the chance of Cloudflare returning
	// an HTML challenge/error page instead of the JSON payload.
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/18.0 Safari/605.1.15")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d from %s: %s", res.StatusCode, url, snippet(body))
	}

	// Guard against HTML error/challenge pages (e.g. Cloudflare) being passed
	// to json.Unmarshal, which otherwise fails with the opaque
	// "invalid character '<' looking for beginning of value".
	contentType := res.Header.Get("Content-Type")
	trimmed := bytes.TrimSpace(body)
	if !strings.Contains(strings.ToLower(contentType), "json") || bytes.HasPrefix(trimmed, []byte("<")) {
		return nil, fmt.Errorf("expected JSON from %s but got content-type %q: %s", url, contentType, snippet(body))
	}

	return body, nil
}

// snippet returns a short, single-line preview of a response body for logging.
func snippet(body []byte) string {
	const max = 120
	s := strings.TrimSpace(string(body))
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > max {
		s = s[:max] + "…"
	}
	return s
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
