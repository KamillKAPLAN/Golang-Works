
// KAMİL KAPLAN

package test

import (
	"encoding/json"
	"github.com/icobani/goweather"
	"io/ioutil"
	"log"
	"testing"
)

func TestForeCast(t *testing.T) {
	apiKey := "Dl8ehKD1lCgL0GT6WCkGfRh9NtSn5GMO"

	c := goweather.NewClient(nil, apiKey)
	var code string = "325236"

	// Daha önce aranan code var ise onu code'a göre dosyaya kaydediyor. Eğer yok ise aranan
	// code.json dosyası oluşturuyor.
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	fileName := code + ".json";
	for _, f := range files {
		if f.Name() == fileName {
			data, err := ioutil.ReadFile(fileName)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(data))
			return
		} else {
			response, _, err := c.Forecast.GetForeCast(code)
			if err != nil {
				log.Fatal(err)
			}
			json_bytelar, _ := json.Marshal(response)
			err2 := ioutil.WriteFile(fileName, json_bytelar, 0777)
			if err2 != nil {
				// print it out
				log.Fatal(err2)
			}
			log.Println(response)
		}
	}
}
