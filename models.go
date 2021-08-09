package coinmarketcap

import (
	"fmt"
	"strings"
)

// QuotesQuery is used to send request parameters to the API for the QuotesLatest call.
type QuotesQuery struct {
	ID          []string
	Slug        []string
	Symbol      []string
	Convert     []string
	ConvertID   []string
	Aux         []string
	SkipInvalid bool
}

// String returns a string representing the URL parameters.
func (q *QuotesQuery) String() string {
	parts := []string{}

	if len(q.ID) > 0 {
		parts = append(parts, fmt.Sprintf("%s=%s", "id", strings.Join(q.ID, ",")))
	}

	if len(q.Slug) > 0 {
		parts = append(parts, fmt.Sprintf("%s=%s", "slug", strings.Join(q.Slug, ",")))
	}

	if len(q.Symbol) > 0 {
		parts = append(parts, fmt.Sprintf("%s=%s", "symbol", strings.Join(q.Symbol, ",")))
	}

	if len(q.Convert) > 0 {
		parts = append(parts, fmt.Sprintf("%s=%s", "convert", strings.Join(q.Convert, ",")))
	}

	if len(q.ConvertID) > 0 {
		parts = append(parts, fmt.Sprintf("%s=%s", "convert_id", strings.Join(q.ConvertID, ",")))
	}

	if len(q.Aux) > 0 {
		parts = append(parts, fmt.Sprintf("%s=%s", "aux", strings.Join(q.Aux, ",")))
	}

	if q.SkipInvalid {
		parts = append(parts, fmt.Sprintf("%s=%s", "skip_invalid", fmt.Sprintf("%v", q.SkipInvalid)))
	}

	return strings.Join(parts, "&")
}

// QuotesResponse is the top level wrapper for a response from QuotesLatest call.
type QuotesResponse struct {
	Status interface{}      `json:"status"`
	Data   map[string]Quote `json:"data"`
}

// Quote represents a quote for a single currency such as BSV, BTC etc.
type Quote struct {
	CirculatingSupply float64                `json:"circulating_supply"`
	CmcRank           int64                  `json:"cmc_rank"`
	DateAdded         string                 `json:"date_added"`
	ID                int64                  `json:"id"`
	IsActive          int64                  `json:"is_active"`
	IsFiat            int64                  `json:"is_fiat"`
	LastUpdated       string                 `json:"last_updated"`
	MaxSupply         int64                  `json:"max_supply"`
	Name              string                 `json:"name"`
	NumMarketPairs    int64                  `json:"num_market_pairs"`
	Quote             map[string]QuoteDetail `json:"quote"`
	Slug              string                 `json:"slug"`
	Symbol            string                 `json:"symbol"`
	Tags              []string               `json:"tags"`
	Platform          interface{}            `json:"platform"`
	TotalSupply       float64                `json:"total_supply"`
}

// QuoteDetail holds the price, change and volume details for the quote.
type QuoteDetail struct {
	LastUpdated      string  `json:"last_updated"`
	MarketCap        float64 `json:"market_cap"`
	PercentChange1h  float64 `json:"percent_change_1h"`
	PercentChange24h float64 `json:"percent_change_24h"`
	PercentChange30d float64 `json:"percent_change_30d"`
	PercentChange60d float64 `json:"percent_change_60d"`
	PercentChange7d  float64 `json:"percent_change_7d"`
	PercentChange90d float64 `json:"percent_change_90d"`
	Price            float64 `json:"price"`
	Volume24h        float64 `json:"volume_24h"`
}
