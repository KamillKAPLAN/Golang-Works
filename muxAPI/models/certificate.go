package models

import (
	"time"
)

type Certificate struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Mail       string    `json:"mail"`
	Name       string    `json:"name"`
	Date       time.Time `json:"date"`
	Type       string    `json:"type"`
	Company    string    `json:"company"`
	EmployeeId uint      `json:"employee_id"`
}
