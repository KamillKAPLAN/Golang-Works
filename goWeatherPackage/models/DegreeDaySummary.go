/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type DegreeDaySummary struct {
	Heating Heating `json:"Heating"`
	Cooling Cooling `json:"Cooling"`
}
