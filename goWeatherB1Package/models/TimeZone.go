/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:45
   Notes   :
*/

package models

import "time"

type TimeZone struct {
	Code             string    `json:"Code,omitempty"`
	Name             string    `json:"Name,omitempty"`
	GmtOffset        float32   `json:"GMTOffset,omitempty"`
	IsDaylightSaving bool      `json:"IsDaylightSaving,omitempty"`
	NextOffsetChange time.Time `json:"NextOffsetChange,omitempty"`
}
