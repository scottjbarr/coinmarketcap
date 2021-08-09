package coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// Host is the host to send requests to.
	Host = "https://pro-api.coinmarketcap.com"

	// HeaderAccept is the "accept" request header.
	HeaderAccept = "accept"

	// ApplicationJSON is the value used for the "accept" request header.
	ApplicationJSON = "application/json"

	// HeaderAPIKey is the request header that holds the user API key.
	HeaderAPIKey = "X-CMC_PRO_API_KEY"
)

// Client is used to communicate with the API.
type Client struct {
	key        string
	Host       string
	HTTPClient *http.Client
}

// New returns a new Client.
func New(key string) (*Client, error) {
	if len(key) == 0 {
		return nil, errors.New("API Key not specified")
	}

	c := Client{
		key:  key,
		Host: Host,
		HTTPClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}

	return &c, nil
}

// QuotesLatest returns the latest market quote for 1 or more cryptocurrencies.
//
// See https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyQuotesLatest
func (c *Client) QuotesLatest(q QuotesQuery) (*QuotesResponse, error) {
	path := fmt.Sprintf("/v1/cryptocurrency/quotes/latest?%s", q.String())

	res := QuotesResponse{}

	if _, err := c.get(path, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// get sends a HTTP GET request to the API.
//
// The in parameter will received the Unmarshalled JSON response if supplied.
func (c Client) get(path string, in interface{}) ([]byte, error) {
	url := c.Host + path

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// set the request headers
	req.Header.Set(HeaderAPIKey, c.key)
	req.Header.Set(HeaderAccept, ApplicationJSON)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if in != nil {
		if err := json.Unmarshal(body, in); err != nil {
			return nil, err
		}
	}

	return body, nil
}
