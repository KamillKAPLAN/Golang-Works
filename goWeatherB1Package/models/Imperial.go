/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:46
   Notes   :
*/

package models

type Imperial struct {
	Value    float64 `json:"Value,omitempty"`
	Unit     string  `json:"Unit,omitempty"`
	UnitType int32   `json:"UnitType,omitempty"`
}
