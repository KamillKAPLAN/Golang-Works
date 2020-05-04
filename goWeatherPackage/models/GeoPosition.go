/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type GeoPosition struct {
	Latitude  float64   `json:"Latitude,omitempty"`
	Longitude float64   `json:"Longitude,omitempty"`
	Elevation Elevation `json:"Elevation"`
}
