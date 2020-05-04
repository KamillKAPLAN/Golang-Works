/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:45
   Notes   :
*/

package models

type Elevation struct {
	Metric   Metric   `json:"Metric"`
	Imperial Imperial `json:"Imperial"`
}
