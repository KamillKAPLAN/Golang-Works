/*
   Name    : Kamil KAPLAN
   Date    : 27.07.2019
*/
package goweather

import (
	"encoding/json"
	"github.com/PROJECTS/goWeatherPackage/models"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"os"
)

// http://dataservice.accuweather.com/locations/v1/cities/search?apikey=...&q=...
func (this *GoWeather) SetLocation(city string, district string) *ErrorStruct {
	var savedLocations []models.Location
	fileIsExist, savedLocations := readLocation()
	// location.json dosyası var mı kontrolü yapıyoruz.
	if fileIsExist {
		for _, value := range savedLocations {
			// dosya var ve aradıımız değerlerle uyuşuyorsa istek atma  içindekileri dön
			if value.AdministrativeArea.EnglishName == city && value.EnglishName == district {
				this.Location = value
				return nil
			}
		}
		return nil
	}
	// eğer dosya yok ise, bu durumda yeni bir istek atıp dosyaya kaydetmemliyiz.
	res, err := resty.R().
		SetQueryParams(map[string]string{
			"apikey": this.GetApiKey(),
			"q":      city + " " + district,
		}).
		SetHeader("Content-Type", "application/json").
		Get(this.LocationURL)

	if err != nil {
		return &ErrorStruct{err.Error()}
	}
	// gönderilen istekte hata yok ise dönen response'ı Unmarshal edip locations'a ekle
	var locations []models.Location
	err = json.Unmarshal(res.Body(), &locations)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}
	// locactions değerini kontrol ediyor. Gerçekten elinde bu struct'ta bir locations var mı diye
	if len(locations) > 0 {
		this.Location = locations[0]
		// dosya kontrolü yapıyor.
		if fileIsExist {
			// Demekki dosya var ama içinde aradığımmız zaman bizim location'ı bulamadık.
			// bu durumda dosyaya yeni bulduğumuz location'e da ekleyebiliriz.
			savedLocations = append(savedLocations, fixEngCharOnLocation(locations[0]))
			errs := writeLocation(savedLocations)
			if errs != nil {
				return &ErrorStruct{errs.Error}
			}

		} else {
			errs := writeLocation(locations)
			if errs != nil {
				return &ErrorStruct{errs.Error}
			}
		}
	} else {
		return &ErrorStruct{"City or distinct name is missing"}
	}
	return nil
}

func readLocation() (bool, []models.Location) {
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

func writeLocation(locations []models.Location) *ErrorStruct {
	var err = os.Remove("locations.json")
	fileBody, _ := json.MarshalIndent(locations, "", " ")
	err = ioutil.WriteFile("locations.json", fileBody, 0644)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}

	return nil
}
