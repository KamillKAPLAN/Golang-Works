/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 16.08.2019 10:15
   Notes   :
*/

package models

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

type LocationTable struct {
	Language                        string    `json:"Language"`
	Version                         int32     `json:"Version,omitempty"`
	Key                             string    `json:"Key,omitempty"`
	Type                            string    `json:"Type,omitempty"`
	Rank                            int32     `json:"Rank,omitempty"`
	LocalizedName                   string    `json:"LocalizedName,omitempty"`
	EnglishName                     string    `json:"EnglishName,omitempty"`
	PrimaryPostalCode               string    `json:"PrimaryPostalCode,omitempty"`
	RegionID                        string    `json:"region_id"`
	RegionName                      string    `json:"region_name"`
	RegionENU                       string    `json:"region_enu"`
	CountryID                       string    `json:"country_id"`
	CountryName                     string    `json:"country_name"`
	CountryENU                      string    `json:"country_enu"`
	AdministrativeAreaID            string    `json:"administrative_area_id"`
	AdministrativeAreaLocalizedName string    `json:"administrative_area_localized_name"`
	AdministrativeAreaNameENU       string    `json:"administrative_area_name_enu"`
	TimeZoneCode                    string    `json:"time_zone_code"`
	TimeZoneName                    string    `json:"time_zone_name"`
	TimeZoneGMTOffset               float32   `json:"time_zone_gmt_offset"`
	TimeZoneNextOffsetChange        time.Time `json:"time_zone_next_offset_change"`
	TimeZoneIsDaylightSaving        bool      `json:"time_zone_is_daylight_saving"`
	GeoPositionLatitude             float64   `json:"geo_position_latitude"`
	GeoPositionLongitude            float64   `json:"geo_position_latitude"`
}

func (this *LocationTable) Save(db *bolt.DB) error {
	// Store the user model in the user bucket using the username as the key.
	var key = fmt.Sprintf("%s|%s|%s", this.EnglishName, this.AdministrativeAreaNameENU, this.Language)
	log.Println(key)
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Locations"))
		if err != nil {
			return err
		}

		encoded, err := json.Marshal(this)
		if err != nil {
			return err
		}
		return b.Put([]byte(key), encoded)
	})
	return err
}

func (this *LocationTable) Get(db *bolt.DB) {
	var key = this.GetPrimaryValue()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Locations"))
		v := b.Get([]byte(key))
		err := json.Unmarshal([]byte(v), &this)
		if err != nil {
			return err
		}
		return nil
	})
}

func (this *LocationTable) Delete(db *bolt.DB) {
	var key = this.GetPrimaryValue()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Locations"))
		err := b.Delete([]byte(key))
		if err != nil {
			return err
		}
		return nil
	})
}

func (this *LocationTable) GetPrimaryValue() string {
	return fmt.Sprintf("%s|%s|%s",
		this.EnglishName,
		this.AdministrativeAreaNameENU,
		this.Language)
}
