/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type Sources struct {
	DataType string `json:"DataType,omitempty"`
	Source   string `json:"Source,omitempty"`
	SourceId int32  `json:"SourceId,omitempty"`
}
