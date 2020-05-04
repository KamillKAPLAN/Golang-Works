package main

import (
	"encoding/json"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"log"
	"time"
)

const (
	baseURL    = "https://ezanvakti.herokuapp.com"
	countryURL = baseURL + "/ulkeler"
	cityURL    = baseURL + "/sehirler"
)

type country struct {
	UlkeAdi   string `json:"UlkeAdi"`
	UlkeAdiEn string `json:"UlkeAdiEn"`
	UlkeID    string `json:"UlkeID"`
}

type cities struct {
	UlkeID string `json:"UlkeID"`
	cities []city `json:"Cities"`
}

type city struct {
	SehirAdi   string `json:"SehirAdi"`
	SehirAdiEn string `json:"SehirAdiEn"`
	SehirID    string `json:"SehirID"`
}

func main() {
	// Resty Default setup
	resty.
		// Set retry count to non zero to enable retries
		SetRetryCount(3).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(5 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(3 * time.Second).
		SetTimeout(1 * time.Minute).
		SetContentLength(true)

	getCountries()
}

func getCountries() {

	res, _ := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(countryURL)

	var country []country
	err := json.Unmarshal(res.Body(), &country)
	if err != nil {
		log.Print("Bam bam bam")
	}

	fileBody, _ := json.MarshalIndent(country, "", " ")

	err = ioutil.WriteFile("country.json", fileBody, 0644)
	if err != nil {
		log.Print("Bam bam bam")
	}

	id := country[186].UlkeID
	getCities(id)
}

func getCities(countryId string) {
	res, _ := resty.R().
		SetQueryParams(map[string]string{
			"ulke": countryId,
		}).
		SetHeader("Content-Type", "application/json").
		Get(cityURL)
	var cities []city
	err := json.Unmarshal(res.Body(), &cities)
	if err != nil {
		log.Print("aaBam bam bam")
	}
	fileBody, _ := json.MarshalIndent(cities, "", " ")

	err = ioutil.WriteFile("city.json", fileBody, 0644)
	if err != nil {
		log.Print("Bam bam bam")
	}
}
