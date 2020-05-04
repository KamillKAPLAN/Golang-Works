package controllers

import (
	"net/http"

	"github.com/ridvansumset/restaurant-menu-app/server/models"

	"github.com/labstack/echo"
)

func ListCategories(c echo.Context) error {
	categories, err := models.ListCategories()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, categories)
}

func GetCategory(c echo.Context) error {
	id := c.Param("category_id")
	category, err := models.GetCategory(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, id+" bulunamadı.")
	}
	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	var category models.Category
	err := c.Bind(&category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createdCategory, err := category.Create() // models'de metot
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, createdCategory)
}

func DeleteCategory(c echo.Context) error {
	id := c.Param("category_id")
	category, err := models.GetCategory(id)
	if err != nil {
		return c.String(http.StatusNotFound, id+"bulunamadı.")
	}
	category.Delete()
	return c.NoContent(http.StatusNoContent)
}

func UpdateCategory(c echo.Context) error {
	id := c.Param("category_id")
	category, getErr := models.GetCategory(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusNotFound, getErr.Error())
	}
	err := c.Bind(&category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, updateErr := category.Update()
	if updateErr != nil {
		return updateErr
	}
	return c.JSON(http.StatusOK, category)
}
