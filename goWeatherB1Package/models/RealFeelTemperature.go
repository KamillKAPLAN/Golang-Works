/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type RealFeelTemperature struct {
	Minimum RFTMinimum `json:"Minimum"`
	Maximum RFTMaximum `json:"Maximum"`
}
