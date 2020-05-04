/*
   Name    : Kamil KAPLAN
   Date    : 27.07.2019
*/
package goweather

import (
	"github.com/PROJECTS/goWeatherPackage/models"
	"os"
	"strings"
)

func fileIsExists(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		// dosyanın ver olup olmasdığını kontrol eeriz.
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

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
