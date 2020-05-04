package models

type School struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Code       int    `json:"code"`
	Name       string `json:"name"`
	City       string `json:"city"`
	Country    string `json:"country"`
	EmployeeId uint   `json:"employee_id"`
}
