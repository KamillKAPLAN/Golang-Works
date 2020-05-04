/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
   Package start
*/
package goweather

import (
	"github.com/PROJECTS/goWeatherPackage/models"
	"gopkg.in/resty.v1"
	"log"
	"time"
)

const (
	baseURL     = "http://dataservice.accuweather.com"
	locationURL = baseURL + "/locations/v1/cities/search"
	forecastURL = baseURL + "/forecasts/v1/daily/1day"
)

// Bir Go Weather struct'ı oluşturdurk. Çünkü AccuWeather API'leri ile iletişimi yönetmek için kullanılacak
type GoWeather struct {
	BaseUrl      string `json:"base_url"`
	LocationURL  string `json:"location_url"`
	ForecastURL  string `json:"forecast_url"`
	ApiKeys      string `json:"api_keys"`
	ApiKeyUsages []models.ApiKeyUsages
	Location     models.Location `json:"location,omitempty"`
	Forecast     models.Forecast `json:"forecast,omitempty"`
	// 'omitempty' kullanmamızdaki temel amaç boş değer gelirse bunları göstermemek amaçlanmıştır.
}

// hata değerlerini dönmesi için bir tane struct yapısı oluşturduk.
type ErrorStruct struct {
	Error string
}

// bu method yeni bir request islemi için kullanılır.
func (this GoWeather) New(ApiKeys string, city string, district string) (*ErrorStruct, *GoWeather) {

	city = fixEnglishChars(city)
	district = fixEnglishChars(district)

	// ApiKeys paramametremizin boş mu olduğunu kontrol ediyoruz.
	if ApiKeys == "" {
		return &ErrorStruct{"ApiKey boş olamaz."}, nil
	}

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

	returVal := &GoWeather{
		BaseUrl:     baseURL,
		LocationURL: locationURL,
		ForecastURL: forecastURL,
		ApiKeys:     ApiKeys,
	}

	// city ve distinct parametresinin boş olmama durumunu kontrol edip işlem yapacağız.
	if city != "" && district != "" {
		log.Print("asdasd")
		err := returVal.SetLocation(city, district)
		returVal.SetForecast()
		if err != nil {
			return err, nil
		}
	}
	return nil, returVal
}
