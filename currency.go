package altinkaynak

import (
	"encoding/xml"
	"time"
)

// currencyResponseEnvelope is the response envelope of the currency service
type currencyResponseEnvelope struct {
	XMLName xml.Name
	Body    currencyResponseBody
}

// currencyResponseBody is the response body of the currency service
type currencyResponseBody struct {
	XMLName             xml.Name
	GetCurrencyResponse currencyResponseData `xml:"GetCurrencyResponse"`
}

// currencyResponseData is the response data of the currency service
type currencyResponseData struct {
	XMLName           xml.Name `xml:"GetCurrencyResponse"`
	GetCurrencyResult string   `xml:"GetCurrencyResult"`
}

// CurrencyService is the service for fetching currency data
type CurrencyService struct {
	apiUrl     string
	payload    string
	currencies map[string]Resource
}

// Get returns a currency by its code
func (cs *CurrencyService) Get(code string) Resource {
	return cs.currencies[code]
}

// Fetch fetches the currency data
func (cs *CurrencyService) Fetch() error {
	response, err := SendRequest("POST", cs.apiUrl, cs.payload)
	if err != nil {
		return err
	}

	currencyResponseEnvelope := &currencyResponseEnvelope{}
	err = xml.Unmarshal(response, currencyResponseEnvelope)
	if err != nil {
		return err
	}

	currencyResultEnvelope := &getResultEnvelope{}
	err = xml.Unmarshal(
		[]byte(currencyResponseEnvelope.Body.GetCurrencyResponse.GetCurrencyResult), currencyResultEnvelope,
	)
	if err != nil {
		return err
	}

	cs.currencies = make(map[string]Resource, len(currencyResultEnvelope.Resources))
	for _, c := range currencyResultEnvelope.Resources {
		c.UpdatedAt, _ = time.ParseInLocation(dateTimeFormat, c.UpdatedAtRaw, location)
		cs.currencies[c.Code] = c
	}

	return nil
}
