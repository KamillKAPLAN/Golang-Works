package goweather

import (
	"log"
	"testing"
)

func TestWeather(t *testing.T) {
	apiKeys := "fAySX5OL8h37IiIt6YKHSUyVQXtph9VM,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs,Z2dcCn8Kr5PDC6Eylj0tRbCSyjrBPlsJ,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs,JLYy1nF8lehLGaGYpdMbgLXAsHgkwHMu"
	var err, goWeather = GoWeather{}.New(apiKeys, "İstanbul", "İncirli")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("GoWeather : %v\n", goWeather)
	}
}
