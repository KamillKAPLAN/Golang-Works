package main

import "time"

type Employee struct {
	ID        uint
	FirstName string
	LastName  string
	StartDate time.Time
	Position  string
	TotalPTO  float32
	Status    string
    TimesOff  []TimeOff
}

type TimeOff struct {
	Type      string
	Amount    float32
	StartDate time.Time
	Status    string
}

var employees = map[string]Employee {
	"962134" : Employee{
		ID:        962134,
		FirstName: "Kamil",
		LastName:  "KAPLAN",
		Position:  "Developer",
		StartDate: time.Now().Add(-13 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
	"176458" : Employee{
		ID:        176158,
		FirstName: "Yasin",
		LastName:  "Memiç",
		Position:  "Software",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Status:    "InActive",
		TotalPTO:  20,
	},
	"160898" : Employee{
		ID:        160898,
		FirstName: "İbrahim",
		LastName:  "Çobani",
		Position:  "CEO",
		StartDate: time.Now().Add(-6 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  40,
	},
	"297365" : Employee{
		ID:        27365,
		FirstName: "Salvo",
		LastName:  "Babani",
		Position:  "Software",
		StartDate: time.Now().Add(-12 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
}

var TimesOff = map[string][]TimeOff{
	"962134" : []TimeOff {
		{
			Type: "Holiday",
			Amount: 8.,
			StartDate: time.Date(2019,07,11,0,0,0,0,time.UTC),
			Status: "Taken",
		}, {
			Type: "PTO",
			Amount: 16.,
			StartDate: time.Date(2019,05,05,0,0,0,0,time.UTC),
			Status: "Scheduled",
		}, {
			Type: "PTO",
			Amount: 16.,
			StartDate: time.Date(2019,07,11,0,0,0,0,time.UTC),
			Status: "Requested",
		},
	},
}