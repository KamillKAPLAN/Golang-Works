package models

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

type Option struct {
	ID         string           `json:"id"`
	CategoryID string           `json:"category_id"`
	ProductID  string           `json:"product_id"`
	Name       string           `json:"name,omitempty"`
	ListOrder  int64            `json:"list_order,omitempty"`
	Required   bool             `json:"required,omitempty"`
	ChoiceType OptionChoiceType `json:"choice_type,omitempty"`
}

type Options []Option

type OptionChoiceType string

const (
	// OptionChoiceTypeSingle should used for single choices
	OptionChoiceTypeSingle OptionChoiceType = "single"
	// OptionChoiceTypeMultiple should used for multiple choices
	OptionChoiceTypeMultiple OptionChoiceType = "multiple"
)

// ListOptions
func ListOptions(categoryID string, productID string) (*Options, error) {
	var newOptionSlice Options
	for _, option := range OptionSlice {
		if option.ProductID == productID {
			// done := make(chan bool)
			// go option.WithChoices(done)
			// <-done
			newOptionSlice = append(newOptionSlice, option)
		}
	}
	return &newOptionSlice, nil
}

// GetOption
func GetOption(optionID string) (*Option, error) {
	for _, option := range OptionSlice {
		if optionID == option.ID {
			// done := make(chan bool)
			// option.WithChoices(done)
			return &option, nil
		}
	}
	return nil, errors.New("not found")
}

// UpdateOptions
func (option *Option) Update() (*Option, error) {
	for i, el := range OptionSlice {
		if option.ID == el.ID {
			OptionSlice = append(OptionSlice[:i], append([]Option{*option}, OptionSlice[i+1:]...)...)
			return option, nil
		}
	}
	return nil, echo.NewHTTPError(http.StatusBadRequest)
}

// func (option *Option) WithChoices(done chan bool) (*Option, error) {
// 	choices, err := ListChoices(option.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	option.Choices = *choices
// 	_, updateErr := option.Update()
// 	if updateErr != nil {
// 		return nil, updateErr
// 	}
// 	done <- true
// 	return option, nil
// }
