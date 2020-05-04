/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type Moon struct {
	Rise      string `json:"Rise,omitempty"`
	EpochRise int64  `json:"EpochRise,omitempty"`
	Set       string `json:"Set,omitempty"`
	EpochSet  int64  `json:"EpochSet,omitempty"`
	Phase     string `json:"Phase,omitempty"`
	Age       int32  `json:"Age,omitempty"`
}
