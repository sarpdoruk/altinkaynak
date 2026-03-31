package altinkaynak

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// GoldService is the service for fetching gold data
type GoldService struct {
	apiUrl string
	golds  map[string]Resource
}

// Get returns a gold resource by its code
func (gs *GoldService) Get(code string) Resource {
	return gs.golds[code]
}

// Fetch fetches the gold data
func (gs *GoldService) Fetch() error {
	response, err := SendRequest("GET", gs.apiUrl)
	if err != nil {
		return err
	}

	var golds []Resource
	err = json.Unmarshal(response, &golds)
	if err != nil {
		return err
	}

	gs.golds = make(map[string]Resource, len(golds))
	for _, c := range golds {
		c.Buy, _ = strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(c.buyString, ".", ""), ",", "."), 64)
		c.Sell, _ = strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(c.sellString, ".", ""), ",", "."), 64)
		c.UpdatedAt, _ = time.ParseInLocation(dateTimeFormat, c.UpdatedAtRaw, location)
		gs.golds[c.Code] = c
	}

	return nil
}
