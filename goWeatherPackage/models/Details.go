/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type Details struct {
	Key                      string    `json:"Key,omitempty"`
	StationCode              string    `json:"StationCode,omitempty"`
	StationGmtOffset         int64     `json:"StationGmtOffset,omitempty"`
	BandMap                  string    `json:"BandMap,omitempty"`
	Climo                    string    `json:"Climo,omitempty"`
	LocalRadar               string    `json:"LocalRadar,omitempty"`
	MediaRegion              string    `json:"MediaRegion,omitempty"`
	Metar                    string    `json:"Metar,omitempty"`
	NXMetro                  string    `json:"NxMetro,omitempty"`
	NXState                  string    `json:"NxState,omitempty"`
	Population               int64     `json:"Population,omitempty"`
	PrimaryWarningCountyCode string    `json:"PrimaryWarningCountyCode,omitempty"`
	PrimaryWarningZoneCode   string    `json:"PrimaryWarningZoneCode,omitempty"`
	Satellite                string    `json:"Satellite,omitempty"`
	Synoptic                 string    `json:"Synoptic,omitempty"`
	MarineStation            string    `json:"MarineStation,omitempty"`
	MarineStationGMTOffset   int64     `json:"MarineStationGmtOffset,omitempty"`
	VideoCode                string    `json:"VideoCode,omitempty"`
	PartnerID                int32     `json:"PartnerId,omitempty"`
	Sources                  []Sources `json:"Sources"`
	CanonicalPostalCode      string    `json:"CanonicalPostalCode,omitempty"`
	CanonicalLocationKey     string    `json:"CanonicalLocationKey,omitempty"`
	LocationStem             string    `json:"LocationStem,omitempty"`
}
