/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:46
   Notes   :
*/

package models

type SupplementalAdminAreas struct {
	Level         int32  `json:"Level,omitempty"`
	LocalizedName string `json:"LocalizedName,omitempty"`
	EnglishName   string `json:"EnglishName,omitempty"`
}
