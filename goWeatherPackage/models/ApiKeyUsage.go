/*
   Name    : Kamil KAPLAN
   Date    : 25.07.2019
*/

package models

type ApiKeyUsages struct {
	ApiKey string `json:"ApiKey"`
	Usage  uint64 `json:"Usage"`
}
