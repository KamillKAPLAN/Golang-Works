package models

import "time"

type Users struct {
	ID        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"primary_key"`
	Phone     string `json:"phone"`
	Age       int    `json:"age"`
	Company   string `json:"company"`
	Active    bool   `json:"active"`
	Address   string `json:"address"`
	Hash      string `json:"hash"`
}

type LoginResp struct {
	TokenVal string    `json:"token_val"`
	Expire   time.Time `json:"expire"`
	Email    string    `json:"email"`
}
