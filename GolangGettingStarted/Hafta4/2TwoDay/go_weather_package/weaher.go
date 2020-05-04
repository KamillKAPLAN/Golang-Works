package goweather

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseURL     = "http://dataservice.accuweather.com"
	locationURL = "http://dataservice.accuweather.com/locations/v1"
	forecastURL = "http://dataservice.accuweather.com/forecasts/v1"
)

// A Client manages communication with the OneSignal API.
type Client struct {
	BaseURL     *url.URL
	LocationURL *url.URL
	ForecastURL *url.URL

	ApiKey string
	Client *http.Client

	Location *LocationService
	Forecast *ForecastService
}

// SuccessResponse  wraps the standard http.Response for several API methods
// that just return a Success flag.
type SuccessResponse struct {
	Success bool `json:"success"`
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Messages []string `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	msg := "Weather returned those error messages:\n - "
	return msg + strings.Join(e.Messages, "\n - ")
}

// NewClient returns a new Weather API client.
func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	locationURL, err := url.Parse(locationURL)
	if err != nil {
		log.Fatal(err)
	}

	forecastURL, err := url.Parse(forecastURL)
	if err != nil {
		log.Fatal(err)
	}

	c := &Client{
		BaseURL:     baseURL,
		LocationURL: locationURL,
		ForecastURL: forecastURL,

		Client: httpClient,
		ApiKey: apiKey,
	}

	c.Location = &LocationService{client: c}
	c.Forecast = &ForecastService{client: c}

	return c
}

// NewRequest creates an API request. path is a relative URL, like "/apps".
// typ is a locationURL, forecastURL or baseURL.
// The AuthKeyType will determine which authorization token (APP or USER) is
// used for the request.
func (c *Client) NewRequest(method string, typ string, path string) (*http.Request, error) {
	// build the URL beging
	wUrl := ""
	switch typ {
	case "location":
		wUrl = c.LocationURL.String()
		break
	case "forecast":
		wUrl = c.ForecastURL.String()
		break
	default:
		wUrl = c.BaseURL.String()
		break
	}
	// build the URL end

	u, err := url.Parse(wUrl + path)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	// create the request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred
func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	// send the request
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	// // log body for debug
	// b := new(bytes.Buffer)
	// b.ReadFrom(resp.Body)
	// log.Println("response body: ", b.String())

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusInternalServerError:
		return &ErrorResponse{
			Messages: []string{"Internal Server Error"},
		}
	default:
		var errResp ErrorResponse
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&errResp)
		if err != nil {
			errResp.Messages = []string{"Couldn't decode response body JSON"}
		}
		return &errResp
	}
}
