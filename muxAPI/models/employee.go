package models

type Employee struct {
	ID           int           `json:"id" gorm:"primary_key"`
	FName        string        `json:"f_name" schema:"f_name"`
	LName        string        `json:"l_name" schema:"l_name,required"`
	Mail         string        `json:"mail" schema:"mail,required"`
	SchoolBegin  uint          `json:"school_begin" schema:"-"`
	SchoolFinish uint          `json:"school_finish" schema:"-"`
	School       []School      `gorm:"foreignkey:employee_id" schema:"-"`
	Certificate  []Certificate `gorm:"foreignkey:employee_id" schema:"-"`
	Education    []Education   `gorm:"foreignkey:employee_id" schema:"-"`
}
