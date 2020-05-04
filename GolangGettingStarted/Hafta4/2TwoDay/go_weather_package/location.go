package goweather

import (
	"net/http"
	"net/url"
	"time"
)

type LocationService struct {
	client *Client
}

type Locations struct {
	Version                int32  `json:"version,omitempty"`
	Key                    string `json:"key,omitempty"`
	Type                   string `json:"type,omitempty"`
	Rank                   int32  `json:"rank,omitempty"`
	LocalizedName          string `json:"LocalizedName,omitempty"`
	EnglishName            string `json:"EnglishName,omitempty"`
	PrimaryPostalCode      string `json:"PrimaryPostalCode,omitempty"`
	Region                 Regions
	Country                Country
	AdministrativeArea     AdministrativeArea
	TimeZone               TimeZone
	GeoPosition            GeoPosition
	IsAlias                bool `json:"ıs_alias,omitempty"`
	SupplementalAdminAreas []SupplementalAdminAreas
	DataSets               []string `json:"data_sets,omitempty"`
	Details                Details
}

type Regions struct {
	ID            string `json:"id,omitempty"`
	LocalizedName string `json:"LocalizedName,omitempty"`
	EnglishName   string `json:"englishName,omitempty"`
}

type Country struct {
	ID            string `json:"id,omitempty"`
	LocalizedName string `json:"localizedName,omitempty"`
	EnglishName   string `json:"englishName,omitempty"`
}

type AdministrativeArea struct {
	ID            string `json:"id,omitempty"`
	LocalizedName string `json:"localizedName,omitempty"`
	EnglishName   string `json:"englishName,omitempty"`
	Level         int32  `json:"level,omitempty"`
	LocalizedType string `json:"localizedType,omitempty"`
	EnglishType   string `json:"englishType,omitempty"`
	CountryID     string `json:"countryId,omitempty"`
}

type TimeZone struct {
	Code             string    `json:"code,omitempty"`
	Name             string    `json:"name,omitempty"`
	GmtOffset        float32   `json:"gmtOffset,omitempty"`
	IsDaylightSaving bool      `json:"ısDaylightSaving,omitempty"`
	NextOffsetChange time.Time `json:"nextOffsetChange,omitempty"`
}

type GeoPosition struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Elevation Elevation
}

type Elevation struct {
	Metric   Metric
	Imperial Imperial
}

type Metric struct {
	Value    float64 `json:"value ,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type Imperial struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type SupplementalAdminAreas struct {
	Level         int32  `json:"level,omitempty"`
	LocalizedName string `json:"localizedName,omitempty"`
	EnglishName   string `json:"englishName,omitempty"`
}

type Details struct {
	Key                      string `json:"key,omitempty"`
	StationCode              string `json:"stationCode,omitempty"`
	StationGmtOffset         int64  `json:"stationGmtOffset,omitempty"`
	BandMap                  string `json:"bandMap,omitempty"`
	Climo                    string `json:"climo,omitempty"`
	LocalRadar               string `json:"localRadar,omitempty"`
	MediaRegion              string `json:"mediaRegion,omitempty"`
	Metar                    string `json:"metar,omitempty"`
	NXMetro                  string `json:"nxMetro,omitempty"`
	NXState                  string `json:"nxState,omitempty"`
	Population               int64  `json:"population,omitempty"`
	PrimaryWarningCountyCode string `json:"primaryWarningCountyCode,omitempty"`
	PrimaryWarningZoneCode   string `json:"primaryWarningZoneCode,omitempty"`
	Satellite                string `json:"satellite,omitempty"`
	Synoptic                 string `json:"synoptic,omitempty"`
	MarineStation            string `json:"marineStation,omitempty"`
	MarineStationGMTOffset   int64  `json:"marineStationGmtOffset,omitempty"`
	VideoCode                string `json:"videoCode,omitempty"`
	PartnerID                int32  `json:"partnerId,omitempty"`
	Sources                  []Sources
	CanonicalPostalCode      string `json:"canonicalPostalCode,omitempty"`
	CanonicalLocationKey     string `json:"canonicalLocationKey,omitempty"`
	LocationStem             string `json:"locationStem,omitempty"`
}

type DMA struct {
	ID          string `json:"id,omitempty"`
	EnglishName string `json:"englishName,omitempty"`
}

type Sources struct {
	DataType string `json:"dataType,omitempty"`
	Source   string `json:"source,omitempty"`
	SourceId int32  `json:"sourceId,omitempty"`
}

// /locations/v1/cities/search?apikey=EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs&q=istanbul
func (l *LocationService) GetCity(city string) (*[]Locations, *http.Response, error) {
	// path := fmt.Sprintf("/cities/search?apikey=%s&q=%s", l.client.ApiKey, city)
	//pull request denemesi
	path := "/cities/search?apikey=" + l.client.ApiKey + "&q=" + city
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}
	req, err := l.client.NewRequest("GET", "location", u.String())
	if err != nil {
		return nil, nil, err
	}

	lResp := new([]Locations)
	resp, err := l.client.Do(req, lResp)
	if err != nil {
		return nil, resp, err
	}
	return lResp, resp, err
}
