/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type Night struct {
	Icon                     int32        `json:"Icon,omitempty"`
	IconPhrase               string       `json:"IconPhrase,omitempty"`
	HasPrecipitation         bool         `json:"HasPrecipitation,omitempty"`
	LocalSource              NLocalSource `json:"LocalSource,omitempty"`
	ShortPhrase              string       `json:"ShortPhrase,omitempty"`
	LongPhrase               string       `json:"LongPhrase,omitempty"`
	PrecipitationProbability int32        `json:"PrecipitationProbability,omitempty"`
	ThunderstormProbability  int32        `json:"ThunderstormProbability,omitempty"`
	RainProbability          int32        `json:"RainProbability,omitempty"`
	SnowProbability          int32        `json:"SnowProbability,omitempty"`
	IceProbability           int32        `json:"IceProbability,omitempty"`
	Wind                     NWind        `json:"Wind,omitempty"`
	WindGust                 NWindGust    `json:"WindGust,omitempty"`
	TotalLiquid              NTotalLiquid `json:"TotalLiquid,omitempty"`
	Rain                     NRain        `json:"Rain,omitempty"`
	Snow                     NSnow        `json:"Snow,omitempty"`
	Ice                      NIce         `json:"Ice,omitempty"`
	HoursOfPrecipitation     float32      `json:"HoursOfPrecipitation,omitempty"`
	HoursOfRain              float32      `json:"HoursOfRain,omitempty"`
	HoursOfSnow              float32      `json:"HoursOfSnow,omitempty"`
	HoursOfIce               float32      `json:"HoursOfIce,omitempty"`
	CloudCover               int32        `json:"CloudCover,omitempty"`
}
