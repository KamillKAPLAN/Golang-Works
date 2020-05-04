/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/
package models

type DailyForecasts struct {
	Date                     string                   `json:"Date,omitempty"`
	EpochDate                int64                    `json:"EpochDate,omitempty"`
	Sun                      Sun                      `json:"Sun"`
	Moon                     Moon                     `json:"Moon"`
	Temperature              Temperature              `json:"Temperature"`
	RealFeelTemperature      RealFeelTemperature      `json:"RealFeelTemperature"`
	RealFeelTemperatureShade RealFeelTemperatureShade `json:"RealFeelTemperatureShade"`
	HoursOfSun               float32                  `json:"HoursOfSun,omitempty"`
	DegreeDaySummary         DegreeDaySummary         `json:"DegreeDaySummary"`
	AirAndPollen             []AirAndPollen           `json:"AirandPollen"`
	Day                      Day                      `json:"Day"`
	Night                    Night                    `json:"Night"`
	Sources                  []string                 `json:"Sources,omitempty"`
	MobieLink                string                   `json:"MobieLink,omitempty"`
	Link                     string                   `json:"Link,omitempty"`
}
