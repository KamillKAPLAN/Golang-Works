/*
   Name    : Kamil KAPLAN
   Date    : 28.07.2019
*/
package goweather

import (
	"encoding/json"
	"github.com/PROJECTS/goWeatherPackage/models"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"os"
)

// http://dataservice.accuweather.com/forecasts/v1/daily/1day/...?apikey=...
func (this *GoWeather) SetForecast() *ErrorStruct {
	fileIsExist := fileIsExists("forecast.json")
	savedForecast := readForecast()
	//dosya varlığını kontrol ederiz.
	if fileIsExist {
		for _, value := range savedForecast {
			if this.Location.Key == value.Key {
				this.Forecast = value
				return nil
			}
		}
	}

	// eğer dosya yoksa burası çalışacak
	res, err := resty.R().
		SetQueryParams(map[string]string{
			"apikey": this.GetApiKey(),
		}).
		SetHeader("Content-Type", "application/json").
		Get(this.ForecastURL + "/" + this.Location.Key)

	if err != nil {
		return &ErrorStruct{err.Error()}
	}

	var forecast models.Forecast
	err = json.Unmarshal(res.Body(), &forecast)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}

	if forecast.Headline.Link != "" {
		forecast.Key = this.Location.Key
		this.Forecast = forecast
		if fileIsExist {
			savedForecast := readForecast()
			if savedForecast == nil {
				return &ErrorStruct{err.Error()}
			}
			savedForecast = append(savedForecast, forecast)
			errs := writeForecast(savedForecast)
			if errs != nil {
				return &ErrorStruct{err.Error()}
			}
		} else {
			var foreCasts []models.Forecast
			foreCasts = append(foreCasts, forecast)
			errs := writeForecast(foreCasts)
			if errs != nil {
				return errs
			}
		}
	} else {
		return &ErrorStruct{err.Error()}
	}
	return nil
}

func readForecast() []models.Forecast {
	var forecast []models.Forecast
	file, _ := ioutil.ReadFile("forecast.json")
	err := json.Unmarshal([]byte(file), &forecast)
	if err != nil {
		return nil
	}
	return forecast
}

func writeForecast(forecast []models.Forecast) *ErrorStruct {
	var err = os.Remove("forecast.json")
	fileBody, _ := json.MarshalIndent(forecast, "", "")
	err = ioutil.WriteFile("forecast.json", fileBody, 0644)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}
	return nil
}
