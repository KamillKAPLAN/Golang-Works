package models

import (
	"time"
)

type Education struct {
	ID             int       `json:"id" gorm:"primary_key"`
	Name           string    `json:"name"`
	DateOfTraining time.Time `json:"dateOfTraining"`
	EmployeeId     uint      `json:"employee_id"`
}
