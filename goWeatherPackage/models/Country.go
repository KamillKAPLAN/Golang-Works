/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type Country struct {
	ID            string `json:"ID,omitempty"`
	LocalizedName string `json:"LocalizedName,omitempty"`
	EnglishName   string `json:"EnglishName,omitempty"`
}
