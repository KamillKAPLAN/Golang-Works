package goweather

import (
	"encoding/json"
	"github.com/icobani/goweather/models"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL     = "http://dataservice.accuweather.com"
	locationURL = baseURL + "/locations/v1/cities/search"
	forecastURL = baseURL + "/forecasts/v1/daily/1day"
)

// A GoWeather manages communication with the AccuWeather API
type GoWeather struct {
	BaseURL      string                `json:"base_url"`
	LocationURL  string                `json:"location_url"`
	ForecastURL  string                `json:"forecast_url"`
	ApiKeys      string                `json:"api_key"`
	Language     string                `json:"language"`
	Details      bool                  `json:"detils"`
	Metric       bool                  `json:"metric"`
	ApiKeyUsages []models.ApiKeyUsages `json:"api_key_usages"`
	Location     models.Location       `json:"location,omitempty"`
	ForeCast     models.Forecast       `json:"fore_cast,omitempty"`
}

type ErrorY struct {
	Error string
}

func (this GoWeather) New(ApiKeys string, city string, district string, language string, details bool, metric bool) (*ErrorY, *GoWeather) {
	// English Characters setting begin
	city = fixEnglishChars(city)
	district = fixEnglishChars(district)
	// English Characters setting end

	// apiKeys controllers begin
	if ApiKeys == "" {
		return &ErrorY{"Api LocationCode is cannot empty"}, nil
	}
	// apiKeys controllers end
	// apiKeys usages beging
	KeyUsages := SetApiKeys(ApiKeys)
	// apiKeys usages end

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
	returnVal := &GoWeather{
		ApiKeys:      ApiKeys,
		ApiKeyUsages: KeyUsages,
		BaseURL:      baseURL,
		LocationURL:  locationURL,
		ForecastURL:  forecastURL,
		Language:     language,
		Details:      details,
		Metric:       metric,
	}

	// TODO:: language olarak default bir dil tanımlanacak mı ??? default dil türkçe olarak ayarladı. sorulacak.
	if city != "" && district != "" {

		err := returnVal.SetLocation(city, district, language)
		if err != nil {
			return err, nil
		}
		err = returnVal.SetForecast()
		if err != nil {
			return err, nil
		}
	}
	return nil, returnVal
}

// English Characters functions begin
func fixEnglishChars(val string) string {

	val = strings.Replace(val, "ğ", "g", 100)
	val = strings.Replace(val, "ü", "u", 100)
	val = strings.Replace(val, "ş", "s", 100)
	val = strings.Replace(val, "ı", "i", 100)
	val = strings.Replace(val, "ö", "o", 100)
	val = strings.Replace(val, "ç", "c", 100)

	val = strings.Replace(val, "Ğ", "G", 100)
	val = strings.Replace(val, "Ü", "U", 100)
	val = strings.Replace(val, "Ş", "S", 100)
	val = strings.Replace(val, "İ", "I", 100)
	val = strings.Replace(val, "Ö", "O", 100)
	val = strings.Replace(val, "Ç", "C", 100)
	return val
}

func fixEngCharOnLocation(val models.Location) models.Location {
	val.EnglishName = fixEnglishChars(val.EnglishName)
	val.AdministrativeArea.EnglishName = fixEnglishChars(val.AdministrativeArea.EnglishName)
	return val
}

// English Characters functions end

// Is there file begin
func fileIsExists(name string) bool {

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Is there file end

// Locations Apı begin
func (this *GoWeather) SetLocation(city string, district string, language string) *ErrorY {

	var savedlocations []models.Location
	fileIsExist, savedlocations := readLocations()
	if fileIsExist {
		// Sorun yok.
		for _, item := range savedlocations {
			if item.AdministrativeArea.EnglishName == city && item.EnglishName == district && item.Language == language {
				this.Location = item
				return nil
			}
		}
	}

	// TODO : Eğer kayıtlarımızda yok ise apiden soracağız.
	res, err := resty.R().
		SetQueryParams(map[string]string{
			"apikey":   this.GetApiKey(),
			"q":        city + " " + district,
			"details":  "true",
			"language": language,
		}).
		SetHeader("Content-Type", "application/json").
		Get(this.LocationURL)
	if err != nil {
		return &ErrorY{err.Error()}
	}

	var locations []models.Location
	err = json.Unmarshal(res.Body(), &locations)
	if err != nil {
		return &ErrorY{err.Error()}
	}

	if len(locations) > 0 {
		locations[0].Language = language
		this.Location = locations[0]
		if fileIsExist {
			// Demekki dosya var ama içinde aradığımmız zaman bizim location'ı bulamadık.
			// bu durumda dosyaya yeni bulduğumuz location'e da ekleyebiliriz.

			savedlocations = append(savedlocations, fixEngCharOnLocation(locations[0]))
			errs := writeLocations(savedlocations)
			if errs != nil {
				return errs
			}
		} else {
			errs := writeLocations(locations)
			if errs != nil {
				return errs
			}
		}

	} else {
		return &ErrorY{"City name is missing"}
	}
	return nil
}

func readLocations() (bool, []models.Location) {
	var locations []models.Location
	var fileIsExist bool
	fileIsExist = fileIsExists("locations.json")
	if fileIsExist {
		file, _ := ioutil.ReadFile("locations.json")
		err := json.Unmarshal([]byte(file), &locations)

		if err != nil {
			return false, nil
		}
		return true, locations
	} else {
		return false, nil
	}
}

func writeLocations(locations []models.Location) *ErrorY {
	var err = os.Remove("locations.json")
	fileBody, _ := json.MarshalIndent(locations, "", " ")
	err = ioutil.WriteFile("locations.json", fileBody, 0644)
	if err != nil {
		return &ErrorY{err.Error()}
	}
	return nil
}

// Locations Apı end

// Forecast Apı begin
func (this *GoWeather) SetForecast() *ErrorY {

	fileIsExist, savedForecast := readForecast()
	// dosya varsa burası çalışacak

	if fileIsExist {
		for i, item := range savedForecast {

			if this.Location.Key == item.LocationCode &&
				this.Language == item.Language {

				if time.Now().Format("2006-01-02") == item.DailyForecasts[0].Date[0:10] {
					this.ForeCast = item
					//log.Print("Kayıt mevcut", this.ForeCast)
					return nil
				} else {
					savedForecast = RemoveIndex(savedForecast, i)
					errs := writeForecasts(savedForecast)
					if errs != nil {
						return errs
					}
					break
				}
			}
		}
	}

	var forecast models.Forecast

	// eğer dosya yoksa burası çalışacak
	res, err := resty.R().
		SetQueryParams(map[string]string{
			"language": this.Language,
			"details":  strconv.FormatBool(this.Details),
			"metric":   strconv.FormatBool(this.Metric),
			"apikey":   this.GetApiKey(),
		}).
		SetHeader("Content-Type", "application/json").
		Get(this.ForecastURL + "/" + this.Location.Key)
	if err != nil {
		return &ErrorY{err.Error()}
	}

	err = json.Unmarshal(res.Body(), &forecast)
	if err != nil {
		return &ErrorY{err.Error()}
	}

	if res.StatusCode() != http.StatusOK {
		type errstr struct {
			Code      string
			Message   string
			Referance string
		}
		var errst errstr
		err = json.Unmarshal(res.Body(), &errst)
		if err != nil {
			return &ErrorY{err.Error()}
		}
		return &ErrorY{errst.Message}
	}

	if forecast.Headline.Link != "" {
		forecast.LocationCode = this.Location.Key
		forecast.Language = this.Language
		this.ForeCast = forecast
		if fileIsExist {
			//var savedForceCast []models.Forecast
			_, savedForceCast := readForecast()
			if savedForceCast == nil {
				return &ErrorY{"Bir hata oluştu."}
			}

			savedForceCast = append(savedForceCast, forecast)
			errs := writeForecasts(savedForceCast)
			if errs != nil {
				return errs
			}
		} else {
			var foreCasts []models.Forecast
			foreCasts = append(foreCasts, forecast)
			errs := writeForecasts(foreCasts)
			if errs != nil {
				return errs
			}
		}
	} else {
		return &ErrorY{"Nedit"}
	}
	return nil
}

func readForecast() (bool, []models.Forecast) {
	var forecast []models.Forecast
	var fileIsExist bool
	fileIsExist = fileIsExists("forecast.json")
	if fileIsExist {
		file, _ := ioutil.ReadFile("forecast.json")
		err := json.Unmarshal([]byte(file), &forecast)

		if err != nil {
			return false, nil
		}
		return true, forecast
	} else {
		return false, nil
	}
}

func writeForecasts(forecast []models.Forecast) *ErrorY {
	var err = os.Remove("forecast.json")
	//var err = os.Remove("forecast.json")
	fileBody, _ := json.MarshalIndent(forecast, "", "")
	err = ioutil.WriteFile("forecast.json", fileBody, 0644)
	//err = ioutil.WriteFile("forecast.json", fileBody, 0644)
	if err != nil {
		return &ErrorY{err.Error()}
	}
	return nil
}

// Forecast Apı begin

// APIKEY's functions bgein
func SetApiKeys(ApiKeys string) []models.ApiKeyUsages {
	var keys []string
	keys = strings.Split(ApiKeys, ",")
	fileIsExist, KeyUsages := readApiKeyUsages()
	if !fileIsExist {
		for _, item := range keys {
			KeyUsages = append(KeyUsages, models.ApiKeyUsages{
				ApiKey: item,
				Usage:  0,
			})
		}
		writeApiKeyUsages(KeyUsages)
	} else {
		var ItemFound bool
		var FileChanged bool
		for _, item := range keys {
			ItemFound = false
			for _, item2 := range KeyUsages {
				if item2.ApiKey == item {
					ItemFound = true
					break
				}
			}
			if !ItemFound {
				KeyUsages = append(KeyUsages, models.ApiKeyUsages{
					ApiKey: item,
					Usage:  0,
				})
				FileChanged = true
			}
		}
		if FileChanged {
			writeApiKeyUsages(KeyUsages)
		}
	}
	return KeyUsages
}

func (this *GoWeather) GetApiKey() string {
	sort.Slice(this.ApiKeyUsages, func(i, j int) bool {
		return this.ApiKeyUsages[i].Usage < this.ApiKeyUsages[j].Usage
	})
	this.ApiKeyUsages[0].Usage += 1
	writeApiKeyUsages(this.ApiKeyUsages)
	return this.ApiKeyUsages[0].ApiKey
}

func readApiKeyUsages() (bool, []models.ApiKeyUsages) {
	var apiKeyUsages []models.ApiKeyUsages
	var fileIsExist bool
	fileIsExist = fileIsExists("apiKeyUsages.json")
	if fileIsExist {
		file, _ := ioutil.ReadFile("apiKeyUsages.json")
		err := json.Unmarshal([]byte(file), &apiKeyUsages)

		if err != nil {
			return false, nil
		}
		return true, apiKeyUsages

	} else {
		return false, nil
	}
}

func writeApiKeyUsages(apiKeyUsages []models.ApiKeyUsages) *ErrorY {
	var err = os.Remove("apiKeyUsages.json")
	fileBody, _ := json.MarshalIndent(apiKeyUsages, "", " ")
	err = ioutil.WriteFile("apiKeyUsages.json", fileBody, 0644)
	if err != nil {
		return &ErrorY{err.Error()}
	}
	return nil
}

// APIKEY'S functions end

func RemoveIndex(s []models.Forecast, index int) []models.Forecast {
	return append(s[:index], s[index+1:]...)
}
