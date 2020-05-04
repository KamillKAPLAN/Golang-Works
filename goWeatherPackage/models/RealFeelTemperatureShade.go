/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type RealFeelTemperatureShade struct {
	Minimum RFTSMinimum `json:"Minimum,omitempty"`
	Maximum RFTSMaximum `json:"Maximum,omitempty"`
}
