/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Sun struct {
	Rise      string `json:"Rise,omitempty"`
	EpochRise int64  `json:"EpochRise,omitempty"`
	Set       string `json:"Set,omitempty"`
	EpochSet  int64  `json:"EpochSet,omitempty"`
}
