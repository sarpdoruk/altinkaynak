package altinkaynak

import (
	"encoding/xml"
	"time"
)

// goldResponseEnvelope is the response envelope of the gold service
type goldResponseEnvelope struct {
	XMLName xml.Name
	Body    goldResponseBody
}

// goldResponseBody is the response body of the gold service
type goldResponseBody struct {
	XMLName         xml.Name
	GetGoldResponse goldResponseData `xml:"GetGoldResponse"`
}

// goldResponseData is the response data of the gold service
type goldResponseData struct {
	XMLName       xml.Name `xml:"GetGoldResponse"`
	GetGoldResult string   `xml:"GetGoldResult"`
}

// GoldService is the service for fetching gold data
type GoldService struct {
	apiUrl     string
	payload    string
	currencies map[string]Resource
}

// Get returns a gold resource by its code
func (cs *GoldService) Get(code string) Resource {
	return cs.currencies[code]
}

// Fetch fetches the gold data
func (cs *GoldService) Fetch() error {
	response, err := SendRequest("POST", cs.apiUrl, cs.payload)
	if err != nil {
		return err
	}

	goldResponseEnvelope := &goldResponseEnvelope{}
	err = xml.Unmarshal(response, goldResponseEnvelope)
	if err != nil {
		return err
	}

	goldResultEnvelope := &getResultEnvelope{}
	err = xml.Unmarshal(
		[]byte(goldResponseEnvelope.Body.GetGoldResponse.GetGoldResult), goldResultEnvelope,
	)
	if err != nil {
		return err
	}

	cs.currencies = make(map[string]Resource, len(goldResultEnvelope.Resources))
	for _, c := range goldResultEnvelope.Resources {
		c.UpdatedAt, _ = time.ParseInLocation(dateTimeFormat, c.UpdatedAtRaw, location)
		cs.currencies[c.Code] = c
	}

	return nil
}
