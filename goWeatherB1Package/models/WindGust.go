/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type WindGust struct {
	Speed     WGSpeed     `json:"Speed"`
	Direction WGDirection `json:"Direction"`
}
