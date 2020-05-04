/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Temperature struct {
	Minimum TMinimum `json:"Minimum"`
	Maximum TMaximum `json:"Maximum"`
}
