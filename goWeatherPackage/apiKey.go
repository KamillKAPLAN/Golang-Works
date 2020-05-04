/*
   Name    : Kamil KAPLAN
   Date    : 28.07.2019
*/
package goweather

import (
	"encoding/json"
	"github.com/PROJECTS/goWeatherPackage/models"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func SetApiKey(ApiKey string) []models.ApiKeyUsages {
	var keys []string
	keys = strings.Split(ApiKey, ",")
	fileIsExist, KeyUsages := readApiKeyUsages()
	if !fileIsExist {
		for _, value := range keys {
			KeyUsages = append(KeyUsages, models.ApiKeyUsages{
				ApiKey: value,
				Usage:  0,
			})
			writeApiKeyUsages(KeyUsages)
		}
	} else {
		var ItemFound bool
		var FileChanged bool
		for _, value := range keys {
			ItemFound = false
			for _, value2 := range KeyUsages {
				if value2.ApiKey == value {
					ItemFound = true
					break
				}
			}
			if !ItemFound {
				KeyUsages = append(KeyUsages, models.ApiKeyUsages{
					ApiKey: value,
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

func writeApiKeyUsages(apiKeyUsages []models.ApiKeyUsages) *ErrorStruct {
	var err = os.Remove("apiKeyUsages.json")
	fileBody, _ := json.MarshalIndent(apiKeyUsages, "", " ")
	err = ioutil.WriteFile("apiKeyUsages.json", fileBody, 0644)
	if err != nil {
		return &ErrorStruct{err.Error()}
	}
	return nil
}
