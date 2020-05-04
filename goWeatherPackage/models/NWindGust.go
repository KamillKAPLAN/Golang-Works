/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type NWindGust struct {
	Speed     NWGSpeed     `json:"Speed,omitempty"`
	Direction NWGDirection `json:"Direction,omitempty"`
}
