package models

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Product structure
type Product struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	ListOrder  int64    `json:"list_order"`
	OptionIDs  []string `json:"option_ids"`
	CategoryID string   `json:"category_id"`
	Price      float64  `json:"price"`
	// CreatedAt      time.Time     `json:"-"`
	// UpdatedAt      time.Time     `json:"-"`
	// DeletedAt      time.Time     `json:"-"`
}

type Products []Product

// ListProducts
func ListProducts(categoryID string) (*Products, error) {
	var newProductSlice Products
	for _, product := range ProductSlice {
		if product.CategoryID == categoryID {
			newProductSlice = append(newProductSlice, product)
		}
	}
	return &newProductSlice, nil
}

// GetProduct
func GetProduct(productID, categoryID string) (*Product, error) {
	for _, product := range ProductSlice {
		if productID == product.ID {
			return &product, nil
		}
	}
	return nil, errors.New("not found")
}

// CreateProduct
func (product *Product) Create() (*Product, error) {
	product.ID = strconv.Itoa(len(ProductSlice) + 1)
	ProductSlice = append(ProductSlice, *product)
	return product, nil
}

// DeleteProduct
func (product *Product) Delete() error {
	for i, produc := range ProductSlice {
		if product.ID == produc.ID {
			ProductSlice = append(ProductSlice[:i], ProductSlice[i+1:]...)
			return nil
		}
	}
	return errors.New("Silinemedi")
}

// UpdateProduct
func (product *Product) Update() (*Product, error) {
	for i, produc := range ProductSlice {
		if product.ID == produc.ID {
			ProductSlice = append(ProductSlice[:i], append([]Product{*product}, ProductSlice[i+1:]...)...)
			return product, nil
		}
	}
	return nil, echo.NewHTTPError(http.StatusBadRequest)
}
