package models

import (
	"errors"
)

type Choice struct {
	ID         string  `json:"id"`
	CategoryID string  `json:"category_id"`
	ProductID  string  `json:"product_id"`
	OptionID   string  `json:"option_id"`
	Name       string  `json:"name,omitempty"`
	ListOrder  int64   `json:"list_order,omitempty"`
	Price      float64 `json:"price,omitempty"`
}

type Choices []Choice

// ListChoices
func ListChoices(categoryID, productID, optionID string) (*Choices, error) {
	var newChoiceSlice Choices
	for _, choice := range ChoiceSlice {
		if choice.OptionID == optionID {
			newChoiceSlice = append(newChoiceSlice, choice)
		}
	}
	// fmt.Println(newChoiceSlice)
	return &newChoiceSlice, nil
}
func GetChoice(choiceID string) (*Choice, error) {
	for _, choice := range ChoiceSlice {
		if choiceID == choice.ID {
			return &choice, nil
		}
	}
	return nil, errors.New("not found")
}
