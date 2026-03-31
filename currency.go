package altinkaynak

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// CurrencyService is the service for fetching currency data
type CurrencyService struct {
	apiUrl     string
	currencies map[string]Resource
}

// Get returns a currency by its code
func (cs *CurrencyService) Get(code string) Resource {
	return cs.currencies[code]
}

// Fetch fetches the currency data
func (cs *CurrencyService) Fetch() error {
	response, err := SendRequest("GET", cs.apiUrl)
	if err != nil {
		return err
	}

	var currencies []Resource
	err = json.Unmarshal(response, &currencies)
	if err != nil {
		return err
	}

	cs.currencies = make(map[string]Resource, len(currencies))
	for _, c := range currencies {
		c.Buy, _ = strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(c.buyString, ".", ""), ",", "."), 64)
		c.Sell, _ = strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(c.sellString, ".", ""), ",", "."), 64)
		c.UpdatedAt, _ = time.ParseInLocation(dateTimeFormat, c.UpdatedAtRaw, location)
		cs.currencies[c.Code] = c
	}

	return nil
}
