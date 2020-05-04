/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Night struct {
	Icon                     int32        `json:"Icon,omitempty"`
	IconPhrase               string       `json:"IconPhrase,omitempty"`
	HasPrecipitation         bool         `json:"HasPrecipitation,omitempty"`
	LocalSource              NLocalSource `json:"LocalSource"`
	ShortPhrase              string       `json:"ShortPhrase,omitempty"`
	LongPhrase               string       `json:"LongPhrase,omitempty"`
	PrecipitationProbability int32        `json:"PrecipitationProbability,omitempty"`
	ThunderstormProbability  int32        `json:"ThunderstormProbability,omitempty"`
	RainProbability          int32        `json:"RainProbability,omitempty"`
	SnowProbability          int32        `json:"SnowProbability,omitempty"`
	IceProbability           int32        `json:"IceProbability,omitempty"`
	Wind                     NWind        `json:"Wind"`
	WindGust                 NWindGust    `json:"WindGust"`
	TotalLiquid              NTotalLiquid `json:"TotalLiquid"`
	Rain                     NRain        `json:"Rain"`
	Snow                     NSnow        `json:"Snow"`
	Ice                      NIce         `json:"Ice"`
	HoursOfPrecipitation     float32      `json:"HoursOfPrecipitation,omitempty"`
	HoursOfRain              float32      `json:"HoursOfRain,omitempty"`
	HoursOfSnow              float32      `json:"HoursOfSnow,omitempty"`
	HoursOfIce               float32      `json:"HoursOfIce,omitempty"`
	CloudCover               int32        `json:"CloudCover,omitempty"`
}
