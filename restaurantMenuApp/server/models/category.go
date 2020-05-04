package models

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Category structure
type Category struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	ListOrder  int64    `json:"list_order"`
	ProductIDs []string `json:"product_ids"`
	// CreatedAt      time.Time     `json:"-"`
	// UpdatedAt      time.Time     `json:"-"`
	// DeletedAt      time.Time     `json:"-"`
}

type Categories []Category

// ListCategory
func ListCategories() (*Categories, error) {
	return &CategorySlice, nil
}

// GetCategory
func GetCategory(id string) (*Category, error) {
	for _, cat := range CategorySlice {
		if id == cat.ID {
			return &cat, nil
		}
	}
	return nil, errors.New("not found")
}

// CreateCategory
func (category *Category) Create() (*Category, error) {
	category.ID = strconv.Itoa(len(CategorySlice) + 1)
	CategorySlice = append(CategorySlice, *category)
	return category, nil
}

// DeleteCategory
func (category *Category) Delete() error {
	for i, cat := range CategorySlice {
		if category.ID == cat.ID {
			CategorySlice = append(CategorySlice[:i], CategorySlice[i+1:]...)
			return nil
		}
	}
	return errors.New("Silinemedi")
}

// UpdateCategory
func (category *Category) Update() (*Category, error) {
	for i, cat := range CategorySlice {
		if category.ID == cat.ID {

			CategorySlice = append(CategorySlice[:i], append([]Category{*category}, CategorySlice[i+1:]...)...)

			return category, nil
		}
	}
	return nil, echo.NewHTTPError(http.StatusBadRequest)
}
