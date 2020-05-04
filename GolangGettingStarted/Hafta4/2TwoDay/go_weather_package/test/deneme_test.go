package test

import (
	"github.com/icobani/goweather"
	"log"
	"testing"
)

func TestDeneme(t *testing.T) {
	apiKey := "EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs"

	c := goweather.NewClient(nil, apiKey)

	log.Println(c)
}
