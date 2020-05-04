/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type Wind struct {
	Speed     WSpeed    `json:"Speed,omitempty"`
	Direction Direction `json:"Direction,omitempty"`
}
