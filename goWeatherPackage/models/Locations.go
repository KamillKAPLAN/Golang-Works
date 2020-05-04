/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type Location struct {
	Version                int32                    `json:"Version,omitempty"`
	Key                    string                   `json:"Key,omitempty"`
	Type                   string                   `json:"Type,omitempty"`
	Rank                   int32                    `json:"Rank,omitempty"`
	LocalizedName          string                   `json:"LocalizedName,omitempty"`
	EnglishName            string                   `json:"EnglishName,omitempty"`
	PrimaryPostalCode      string                   `json:"PrimaryPostalCode,omitempty"`
	Region                 Regions                  `json:"Region"`
	Country                Country                  `json:"Country"`
	AdministrativeArea     AdministrativeArea       `json:"AdministrativeArea"`
	TimeZone               TimeZone                 `json:"TimeZone"`
	GeoPosition            GeoPosition              `json:"GeoPosition"`
	IsAlias                bool                     `json:"IsAlias,omitempty"`
	SupplementalAdminAreas []SupplementalAdminAreas `json:"SupplementalAdminAreas"`
	DataSets               []string                 `json:"DataSets,omitempty"`
	Details                Details                  `json:"Details"`
}
