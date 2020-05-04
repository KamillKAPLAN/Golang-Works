package models

import "time"

// Employee TABLE
type Employee struct {
	ID             int       `json:"id" gorm:"primary_key"`
	EmployeeName   string    `json:"employee_name"`
	EmployeeSalary int       `json:"employee_salary"`
	EmployeeAge    uint8     `json:"employee_age"`
	CreatedDate    time.Time `json:"created_date"`
}
