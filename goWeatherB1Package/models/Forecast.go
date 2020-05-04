/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Forecast struct {
	Language       string           `json:"Language"`
	LocationCode   string           `json:"LocationCode"`
	Headline       Headline         `json:"Headline"`
	DailyForecasts []DailyForecasts `json:"DailyForecasts"`
}
