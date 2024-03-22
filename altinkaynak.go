// Package altinkaynak provides services for fetching currency and gold data from [Altinkaynak API].
//
// [Altinkaynak API]: https://www.altinkaynak.com/Araclar/Servisler
package altinkaynak

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	dateTimeFormat  = "02.1.2006 15:04:05"
	payloadTemplate = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Header>
    <AuthHeader xmlns="http://data.altinkaynak.com/">
      <Username>%s</Username>
      <Password>%s</Password>
    </AuthHeader>
  </soap:Header>
  <soap:Body>
    <%s xmlns="http://data.altinkaynak.com/" />
  </soap:Body>
</soap:Envelope>`
)

// Currency constants
const (
	CURRENCY_USD string = "USD"
	CURRENCY_EUR        = "EUR"
	CURRENCY_CHF        = "CHF"
	CURRENCY_GBP        = "GBP"
	CURRENCY_JPY        = "JPY"
	CURRENCY_SAR        = "SAR"
	CURRENCY_AUD        = "AUD"
	CURRENCY_CAD        = "CAD"
	CURRENCY_RUB        = "RUB"
	CURRENCY_AZN        = "AZN"
	CURRENCY_CNY        = "CNY"
	CURRENCY_RON        = "RON"
	CURRENCY_AED        = "AED"
	CURRENCY_BGN        = "BGN"
	CURRENCY_KWD        = "KWD"
)

// Gold constants
const (
	GOLD_HAS_TOPTAN      string = "HH_T"
	GOLD_KULCE_TOPTAN           = "CH_T"
	GOLD_GRAM_TOPTAN            = "GAT"
	GOLD_22_AYAR_HURDA          = "B_T"
	GOLD_ATA_TOPTAN             = "A_T"
	GOLD_ESKI_CEYREK            = "EC"
	GOLD_ESKI_YARIM             = "EY"
	GOLD_ESKI_TEKLIK            = "ET"
	GOLD_ESKI_GREMSE            = "EG"
	GOLD_GUMUS                  = "AG_T"
	GOLD_ATA_CUMHURIYET         = "A"
	GOLD_22_AYAR_BILEZIK        = "B"
	GOLD_18_AYAR                = "18"
	GOLD_14_AYAR                = "14"
	GOLD_CEYREK                 = "C"
	GOLD_YARIM                  = "Y"
	GOLD_TEKLIK                 = "T"
	GOLD_GREMSE                 = "G"
	GOLD_ATA_BESLI              = "A5"
	GOLD_RESAT                  = "R"
	GOLD_HAMIT                  = "H"
	GOLD_GRAM                   = "GA"
	GOLD_ONS                    = "ONS"
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
	apiUrl   string
	username string
	password string
}

// getResultEnvelope represents the result envelope of the Altinkaynak API
type getResultEnvelope struct {
	XMLName   xml.Name   `xml:"Kurlar"`
	Resources []Resource `xml:"Kur"`
}

// Option represents a functional option for the Altinkaynak API
type Option func(*config)

// Altinkaynak represents the Altinkaynak API
type Altinkaynak struct {
	CurrencyService *CurrencyService
	GoldService     *GoldService
	MainService     *MainService
}

// Resource represents a currency or gold resource
type Resource struct {
	Code         string  `xml:"Kod"`
	Name         string  `xml:"Aciklama"`
	Buy          float64 `xml:"Alis"`
	Sell         float64 `xml:"Satis"`
	UpdatedAtRaw string  `xml:"GuncellenmeZamani"`
	UpdatedAt    time.Time
}

// SendRequest sends a POST request to the given URL with the given payload
func SendRequest(method string, url string, payload string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	return body, nil
}

// WithCredentials sets the username and password for the Altinkaynak API
func WithCredentials(username string, password string) Option {
	return func(c *config) {
		c.username = username
		c.password = password
	}
}

func WithApiUrl(url string) Option {
	return func(c *config) {
		c.apiUrl = url
	}
}

func NewAltinkaynak(opts ...Option) *Altinkaynak {
	// Set location
	location, _ = time.LoadLocation("Europe/Istanbul")

	// Default credentials
	config := &config{
		apiUrl:   "http://data.altinkaynak.com/DataService.asmx",
		username: "AltinkaynakWebServis",
		password: "AltinkaynakWebServis",
	}

	// Apply options
	for _, opt := range opts {
		opt(config)
	}

	// Return Altinkaynak instance
	return &Altinkaynak{
		CurrencyService: &CurrencyService{
			apiUrl:  config.apiUrl,
			payload: fmt.Sprintf(payloadTemplate, config.username, config.password, "GetCurrency"),
		},
		GoldService: &GoldService{
			apiUrl:  config.apiUrl,
			payload: fmt.Sprintf(payloadTemplate, config.username, config.password, "GetGold"),
		},
		MainService: &MainService{
			apiUrl:  config.apiUrl,
			payload: fmt.Sprintf(payloadTemplate, config.username, config.password, "GetMain"),
		},
	}
}
