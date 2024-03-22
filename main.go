package altinkaynak

import (
	"encoding/xml"
	"time"
)

// mainResponseEnvelope is the response envelope of the main service
type mainResponseEnvelope struct {
	XMLName xml.Name
	Body    mainResponseBody
}

// mainResponseBody is the response body of the main service
type mainResponseBody struct {
	XMLName         xml.Name
	GetMainResponse mainResponseData `xml:"GetMainResponse"`
}

// mainResponseData is the response data of the main service
type mainResponseData struct {
	XMLName       xml.Name `xml:"GetMainResponse"`
	GetMainResult string   `xml:"GetMainResult"`
}

// MainService is the service for fetching main data
type MainService struct {
	apiUrl     string
	payload    string
	currencies map[string]Resource
}

// Get returns a main resource by its code
func (cs *MainService) Get(code string) Resource {
	return cs.currencies[code]
}

// Fetch fetches the main data
func (cs *MainService) Fetch() error {
	response, err := SendRequest("POST", cs.apiUrl, cs.payload)
	if err != nil {
		return err
	}

	mainResponseEnvelope := &mainResponseEnvelope{}
	err = xml.Unmarshal(response, mainResponseEnvelope)
	if err != nil {
		return err
	}

	mainResultEnvelope := &getResultEnvelope{}
	err = xml.Unmarshal(
		[]byte(mainResponseEnvelope.Body.GetMainResponse.GetMainResult), mainResultEnvelope,
	)
	if err != nil {
		return err
	}

	cs.currencies = make(map[string]Resource, len(mainResultEnvelope.Resources))
	for _, c := range mainResultEnvelope.Resources {
		c.UpdatedAt, _ = time.ParseInLocation(dateTimeFormat, c.UpdatedAtRaw, location)
		cs.currencies[c.Code] = c
	}

	return nil
}
