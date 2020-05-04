/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:47
   Notes   :
*/

package models

type Sources struct {
	DataType string `json:"DataType,omitempty"`
	Source   string `json:"Source,omitempty"`
	SourceId int32  `json:"SourceId,omitempty"`
}
