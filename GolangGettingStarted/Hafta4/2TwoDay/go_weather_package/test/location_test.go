package test

import (
	"github.com/icobani/goweather"
	"log"
	"testing"
)

func TestLocation(t *testing.T) {
	apiKey := "Dl8ehKD1lCgL0GT6WCkGfRh9NtSn5GMO"

	c := goweather.NewClient(nil, apiKey)

	response, _, err := c.Location.GetCity("istanbul")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
