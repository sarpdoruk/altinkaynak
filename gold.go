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
	for _, g := range golds {
		g.Buy, _ = strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(g.BuyString, ".", ""), ",", "."), 64)
		g.Sell, _ = strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(g.SellString, ".", ""), ",", "."), 64)
		g.UpdatedAt, _ = time.ParseInLocation(dateTimeFormat, g.UpdatedAtRaw, location)
		gs.golds[g.Code] = g
	}

	return nil
}
