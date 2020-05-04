package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Author struct {
	FirstName string
	LastName  string
	Email     string
}

type Example struct {
	Id          int64
	Title       string
	Slug        string
	DateCreated time.Time
	DateUpdated time.Time
	Authors     []Author
}

// MarshalBinary -
func (e *Example) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Example) UnmarshalBinary(data []byte) (error) {
	err := json.Unmarshal(data, &e);
	if err != nil {
		return err
	}
	return nil
}

func main() {
	example := Example{
		Id:          12345,
		Title:       "Example Title",
		Slug:        "Example Slug",
		DateCreated: time.Now().UTC(),
		Authors: []Author{
			Author{
				FirstName: "Nevio",
				LastName:  "Vesic",
				Email:     "nevio.vesic@gmail.com",
			},
		},
	}
	eb, eberr := example.MarshalBinary()
	if eberr != nil {
		fmt.Errorf("Unable to marshal example struct due to: %s \n", eberr)
	}

	fmt.Printf("Example.MarshalBinary: %v \n", eb)

	var newExample Example

	if err := newExample.UnmarshalBinary(eb); err != nil {
		fmt.Errorf("Unable to unmarshal data into the new example struct due to: %s \n", err)
	}

	fmt.Printf("Example.UnmarshalBinary: %+v \n", newExample)
}
