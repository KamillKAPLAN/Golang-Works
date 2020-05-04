/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type DLocalSource struct {
	Id          int32  `json:"Id,omitempty"`
	Name        string `json:"Name,omitempty"`
	WeatherCode string `json:"WeatherCode,omitempty"`
}
