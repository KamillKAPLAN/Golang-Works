/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type AdministrativeArea struct {
	ID            string `json:"ID,omitempty"`
	LocalizedName string `json:"LocalizedName,omitempty"`
	EnglishName   string `json:"EnglishName,omitempty"`
	Level         int32  `json:"Level,omitempty"`
	LocalizedType string `json:"LocalizedType,omitempty"`
	EnglishType   string `json:"EnglishType,omitempty"`
	CountryID     string `json:"CountryID,omitempty"`
}
