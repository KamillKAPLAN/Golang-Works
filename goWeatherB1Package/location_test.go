package goweather

import (
	"github.com/icobani/goweather/models"
	bolt "go.etcd.io/bbolt"
	"log"
	"testing"
	"time"
)

func TestLocation(t *testing.T) {

	apiKeys := "Z2dcCn8Kr5PDC6Eylj0tRbCSyjrBPlsJ,JbuSGqPN2WrQvChrPxx8vXjd1ekYshQh,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs0,sJXIyoAmMS68m2ACUWuAxGnzEp0OGcv0,fAySX5OL8h37IiIt6YKHSUyVQXtph9VM,JLYy1nF8lehLGaGYpdMbgLXAsHgkwHMu,yLxyo5AWCoO8vHOLphOty0QsFxj5sUe2,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs,tdfew2mhPXGw3rs18ACrS1DHk5mMGg49"
	var err, goWeather = GoWeather{}.New(apiKeys, "Londra", "LND,GB", "tr-TR", true, true)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Location : %v\nGoWeather : %v\n", goWeather.Location, goWeather.ForeCast)
	}

}

func TestSetupBoltDB(t *testing.T) {
	//init db
	db, err := bolt.Open("weather.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return err
		}

		_, err = root.CreateBucketIfNotExists([]byte("Location"))
		if err != nil {
			return err
		}
		log.Println("DB Setup Done")
		return nil
	})
}
func TestSaveLocation(t *testing.T) {
	apiKeys := "Z2dcCn8Kr5PDC6Eylj0tRbCSyjrBPlsJ,JbuSGqPN2WrQvChrPxx8vXjd1ekYshQh,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs0,sJXIyoAmMS68m2ACUWuAxGnzEp0OGcv0,fAySX5OL8h37IiIt6YKHSUyVQXtph9VM,JLYy1nF8lehLGaGYpdMbgLXAsHgkwHMu,yLxyo5AWCoO8vHOLphOty0QsFxj5sUe2,EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs,tdfew2mhPXGw3rs18ACrS1DHk5mMGg49"
	var err, Weather = GoWeather{}.New(apiKeys, "İstanbul", "Yenibosna", "tr-tr", true, true)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("GoWeather : %v\n", Weather.ForeCast.DailyForecasts[0].Date)
	}

	db, errDb := bolt.Open("weather.db", 0600, nil)
	if errDb != nil {
		log.Fatal(err)
	}
	var locationTable = models.LocationTable{
		Language:                        Weather.Location.Language,
		Version:                         Weather.Location.Version,
		Key:                             Weather.Location.Key,
		Type:                            Weather.Location.Type,
		Rank:                            Weather.Location.Rank,
		LocalizedName:                   Weather.Location.LocalizedName,
		EnglishName:                     Weather.Location.EnglishName,
		PrimaryPostalCode:               Weather.Location.PrimaryPostalCode,
		RegionID:                        Weather.Location.Region.ID,
		RegionName:                      Weather.Location.Region.LocalizedName,
		RegionENU:                       Weather.Location.Region.EnglishName,
		CountryID:                       Weather.Location.Country.ID,
		CountryName:                     Weather.Location.Country.LocalizedName,
		CountryENU:                      Weather.Location.Country.EnglishName,
		AdministrativeAreaID:            Weather.Location.AdministrativeArea.ID,
		AdministrativeAreaLocalizedName: Weather.Location.AdministrativeArea.LocalizedName,
		AdministrativeAreaNameENU:       Weather.Location.AdministrativeArea.EnglishName,
		TimeZoneCode:                    Weather.Location.TimeZone.Code,
		TimeZoneName:                    Weather.Location.TimeZone.Name,
		TimeZoneNextOffsetChange:        Weather.Location.TimeZone.NextOffsetChange,
		TimeZoneGMTOffset:               Weather.Location.TimeZone.GmtOffset,
		TimeZoneIsDaylightSaving:        Weather.Location.TimeZone.IsDaylightSaving,
	}
	locationTable.Save(db)
	locationTable.Get(db)
	db.Close()
}
func TestGetLocation(t *testing.T) {
	var locationTable = models.LocationTable{
		EnglishName:               "Yenibosna",
		AdministrativeAreaNameENU: "Istanbul",
		Language:                  "tr-tr",
	}
	db, errDb := bolt.Open("weather.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if errDb != nil {
		log.Fatal(errDb)
	}
	defer db.Close()
	locationTable.Get(db)
	log.Printf("\nİlçe : %s \n il : %s", locationTable.LocalizedName, locationTable.AdministrativeAreaLocalizedName)
}
func TestDeleteLocation(t *testing.T) {

	db, errDb := bolt.Open("weather.db", 0600, nil)
	if errDb != nil {
		log.Fatal(errDb)
	}
	var locationTable = models.LocationTable{
		Language:                  "tr-tr",
		EnglishName:               "1",
		AdministrativeAreaNameENU: "2",
		LocalizedName:             "Piti",
	}
	locationTable.Save(db)
	locationTable.Get(db)
	log.Printf("%s", locationTable.LocalizedName)
	locationTable.Delete(db)
	locationTable.Get(db)
	log.Printf("%s", locationTable.LocalizedName)

}
