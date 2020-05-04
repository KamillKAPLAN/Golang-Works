/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type NWind struct {
	Speed     WNSpeed    `json:"Speed",omitempty`
	Direction NDirection `json:"Direction,omitempty"`
}
