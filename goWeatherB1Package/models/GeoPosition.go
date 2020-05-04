/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.07.2019 11:45
   Notes   :
*/

package models

type GeoPosition struct {
	Latitude  float64   `json:"Latitude,omitempty"`
	Longitude float64   `json:"Longitude,omitempty"`
	Elevation Elevation `json:"Elevation"`
}
