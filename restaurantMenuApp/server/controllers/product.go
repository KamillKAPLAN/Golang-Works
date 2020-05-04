package controllers

import (
	"errors"
	"net/http"

	"github.com/ridvansumset/restaurant-menu-app/server/models"

	"github.com/labstack/echo"
)

func ListProducts(c echo.Context) error {
	categoryID := c.Param("category_id")
	products, err := models.ListProducts(categoryID)
	if err != nil {
		return errors.New("Ürün bulunamadı")
	}
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	categoryID := c.Param("category_id")
	productID := c.Param("product_id")
	product, err := models.GetProduct(productID, categoryID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, productID+" bulunamadı.")
	}
	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	categoryID := c.Param("category_id")
	var product models.Product
	err := c.Bind(&product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	product.CategoryID = categoryID

	createdProduct, err := product.Create() // models'de metot
	if err != nil {
		return err
	}
	category, err := models.GetCategory(categoryID)
	category.ProductIDs = append(category.ProductIDs, product.ID)
	if err != nil {
		return err
	}
	_, updateErr := category.Update()
	if updateErr != nil {
		return updateErr
	}
	return c.JSON(http.StatusCreated, createdProduct)
}

// func DeleteProduct(c echo.Context) error {
// 	id := c.Param("product_id")
// 	product, err := models.GetProduct(id)
// 	if err != nil {
// 		return c.String(http.StatusNotFound, id+"bulunamadı.")
// 	}
// 	product.Delete()
// 	return c.NoContent(http.StatusNoContent)
// }
//
// func UpdateProduct(c echo.Context) error {
// 	id := c.Param("product_id")
// 	product, getErr := models.GetProduct(id)
// 	if getErr != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, getErr.Error())
// 	}
//
// 	err := c.Bind(&product)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	_, updateErr := product.Update()
// 	if updateErr != nil {
// 		return updateErr
// 	}
// 	return c.JSON(http.StatusOK, product)
// }
