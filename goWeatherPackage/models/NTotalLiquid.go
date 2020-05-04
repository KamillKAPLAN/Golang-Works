/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type NTotalLiquid struct {
	Value    float64 `json:"Value,omitempty"`
	Unit     string  `json:"Unit,omitempty"`
	UnitType int32   `json:"UnitType,omitempty"`
}
