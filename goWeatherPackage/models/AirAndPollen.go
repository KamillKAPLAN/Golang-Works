/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type AirAndPollen struct {
	Name          string `json:"Name,omitempty"`
	Value         int32  `json:"Value,omitempty"`
	Category      string `json:"Category,omitempty"`
	CategoryValue int32  `json:"CategoryValue,omitempty"`
	Type          string `json:"Type,omitempty"`
}
