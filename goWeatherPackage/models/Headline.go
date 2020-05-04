/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type Headline struct {
	EffectiveDate      string `json:"EffectiveDate,omitempty"`
	EffectiveEpochDate int64  `json:"EffectiveEpochDate,omitempty"`
	Severity           int32  `json:"Severity,omitempty"`
	Text               string `json:"Text,omitempty"`
	Category           string `json:"Category,omitempty"`
	EndDate            string `json:"EndDate,omitempty"`
	EndEpochDate       int64  `json:"EndEpochDate,omitempty"`
	MobileLink         string `json:"MobileLink,omitempty"`
	Link               string `json:"Link,omitempty"`
}