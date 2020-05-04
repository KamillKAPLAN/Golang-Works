package goweather

import (
	"net/http"
	"net/url"
)

type ForecastService struct {
	client *Client
}

type Forecast struct {
	Headline       Headline
	DailyForecasts []DailyForecasts
}

type Headline struct {
	EffectiveDate      string `json:"effectiveDate,omitempty"`
	EffectiveEpochDate int64  `json:"effectiveEpochDate,omitempty"`
	Severity           int32  `json:"severity,omitempty"`
	Text               string `json:"text,omitempty"`
	Category           string `json:"category,omitempty"`
	EndDate            string `json:"endDate,omitempty"`
	EndEpochDate       int64  `json:"endEpochDate,omitempty"`
	MobileLink         string `json:"mobileLink,omitempty"`
	Link               string `json:"link,omitempty"`
}

type DailyForecasts struct {
	Date                     string `json:"date,omitempty"`
	EpochDate                int64  `json:"epochDate,omitempty"`
	Sun                      Sun
	Moon                     Moon
	Temperature              Temperature
	RealFeelTemperature      RealFeelTemperature
	RealFeelTemperatureShade RealFeelTemperatureShade
	HoursOfSun               float32 `json:"hoursOfSun,omitempty"`
	DegreeDaySummary         DegreeDaySummary
	AirAndPollen             []AirAndPollen
	Day                      Day
	Night                    Night
	Sources                  []string `json:"sources,omitempty"`
	MobieLink                string   `json:"mobieLink,omitempty"`
	Link                     string   `json:"link,omitempty"`
}

type Sun struct {
	Rise      string `json:"rise,omitempty"`
	EpochRise int64  `json:"epochRise,omitempty"`
	Set       string `json:"set,omitempty"`
	EpochSet  int64  `json:"epochSet,omitempty"`
}

type Moon struct {
	Rise      string `json:"rise,omitempty"`
	EpochRise int64  `json:"epochRise,omitempty"`
	Set       string `json:"set,omitempty"`
	EpochSet  int64  `json:"epochSet,omitempty"`
	Phase     string `json:"phase,omitempty"`
	Age       int32  `json:"age,omitempty"`
}

type Temperature struct {
	Minimum TMinimum
	Maximum TMaximum
}

type TMinimum struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type TMaximum struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type RealFeelTemperature struct {
	Minimum RFTMinimum
	Maximum RFTMaximum
}

type RFTMinimum struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type RFTMaximum struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type RealFeelTemperatureShade struct {
	Minimum RFTSMinimum
	Maximum RFTSMaximum
}

type RFTSMinimum struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type RFTSMaximum struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type DegreeDaySummary struct {
	Heating Heating
	Cooling Cooling
}

type Heating struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type Cooling struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type AirAndPollen struct {
	Name          string `json:"name,omitempty"`
	Value         int32  `json:"value,omitempty"`
	Category      string `json:"category,omitempty"`
	CategoryValue int32  `json:"categoryValue,omitempty"`
	Type          string `json:"type,omitempty"`
}

type Day struct {
	Icon                     int32  `json:"ıcon,omitempty"`
	IconPhrase               string `json:"ıconPhrase,omitempty"`
	HasPrecipitation         bool   `json:"hasPrecipitation,omitempty"`
	LocalSource              DLocalSource
	ShortPhrase              string `json:"shortPhrase,omitempty"`
	LongPhrase               string `json:"longPhrase,omitempty"`
	PrecipitationProbability int32  `json:"precipitationProbability,omitempty"`
	ThunderstormProbability  int32  `json:"thunderstormProbability,omitempty"`
	RainProbability          int32  `json:"rainProbability,omitempty"`
	SnowProbability          int32  `json:"snowProbability,omitempty"`
	IceProbability           int32  `json:"ıceProbability,omitempty"`
	Wind                     Wind
	WindGust                 WindGust
	TotalLiquid              TotalLiquid
	Rain                     Rain
	Snow                     Snow
	Ice                      Ice
	HoursOfPrecipitation     float32 `json:"hoursOfPrecipitation,omitempty"`
	HoursOfRain              float32 `json:"hoursOfRain,omitempty"`
	HoursOfSnow              float32 `json:"hoursOfSnow,omitempty"`
	HoursOfIce               float32 `json:"hoursOfIce,omitempty"`
	CloudCover               int32   `json:"cloudCover,omitempty"`
}

type DLocalSource struct {
	Id          int32  `json:"ıd,omitempty"`
	Name        string `json:"name,omitempty"`
	WeatherCode string `json:"weatherCode,omitempty"`
}

type Wind struct {
	Speed     WSpeed
	Direction Direction
}

type WSpeed struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type Direction struct {
	Degress   float64 `json:"degress,omitempty"`
	Localized string  `json:"localized,omitempty"`
	English   string  `json:"english,omitempty"`
}

type WindGust struct {
	Speed     WGSpeed
	Direction WGDirection
}

type WGSpeed struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type WGDirection struct {
	Degress   float64 `json:"degress,omitempty"`
	Localized string  `json:"localized,omitempty"`
	English   string  `json:"english,omitempty"`
}

type TotalLiquid struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type Rain struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type Snow struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type Ice struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type Night struct {
	Icon                     int32  `json:"ıcon,omitempty"`
	IconPhrase               string `json:"ıconPhrase,omitempty"`
	HasPrecipitation         bool   `json:"hasPrecipitation,omitempty"`
	LocalSource              NLocalSource
	ShortPhrase              string `json:"shortPhrase,omitempty"`
	LongPhrase               string `json:"longPhrase,omitempty"`
	PrecipitationProbability int32  `json:"precipitationProbability,omitempty"`
	ThunderstormProbability  int32  `json:"thunderstormProbability,omitempty"`
	RainProbability          int32  `json:"rainProbability,omitempty"`
	SnowProbability          int32  `json:"snowProbability,omitempty"`
	IceProbability           int32  `json:"ıceProbability,omitempty"`
	Wind                     NWind
	WindGust                 NWindGust
	TotalLiquid              NTotalLiquid
	Rain                     NRain
	Snow                     NSnow
	Ice                      NIce
	HoursOfPrecipitation     float32 `json:"hoursOfPrecipitation,omitempty"`
	HoursOfRain              float32 `json:"hoursOfRain,omitempty"`
	HoursOfSnow              float32 `json:"hoursOfSnow,omitempty"`
	HoursOfIce               float32 `json:"hoursOfIce,omitempty"`
	CloudCover               int32   `json:"cloudCover,omitempty"`
}

type NLocalSource struct {
	Id          int32  `json:"ıd,omitempty"`
	Name        string `json:"name,omitempty"`
	WeatherCode string `json:"weatherCode,omitempty"`
}

type NWind struct {
	Speed     WNSpeed
	Direction NDirection
}

type WNSpeed struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type NDirection struct {
	Degress   float64 `json:"degress,omitempty"`
	Localized string  `json:"localized,omitempty"`
	English   string  `json:"english,omitempty"`
}

type NWindGust struct {
	Speed     NWGSpeed
	Direction NWGDirection
}

type NWGSpeed struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type NWGDirection struct {
	Degress   float64 `json:"degress,omitempty"`
	Localized string  `json:"localized,omitempty"`
	English   string  `json:"english,omitempty"`
}

type NTotalLiquid struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type NRain struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type NSnow struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

type NIce struct {
	Value    float64 `json:"value,omitempty"`
	Unit     string  `json:"unit,omitempty"`
	UnitType int32   `json:"unitType,omitempty"`
}

// /forecasts/v1/daily/1day/318251?apikey=EUCeE3zp9BHDZLQDkBU7Y6KvuI3HPozs
func (f *ForecastService) GetForeCast(code string) (*Forecast, *http.Response, error) {
	path := "/daily/1day/" + code + "?apikey=" + f.client.ApiKey
	u, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}
	req, err := f.client.NewRequest("GET", "forecast", u.String())
	if err != nil {
		return nil, nil, err
	}

	fResp := new(Forecast)
	resp, err := f.client.Do(req, fResp)
	if err != nil {
		return nil, resp, err
	}
	return fResp, resp, err
}