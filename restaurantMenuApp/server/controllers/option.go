package controllers

import (
	"net/http"

	"github.com/ridvansumset/restaurant-menu-app/server/models"

	"github.com/labstack/echo"
)

func ListOptions(c echo.Context) error {
	categoryID := c.Param("category_id")
	productID := c.Param("product_id")
	options, err := models.ListOptions(categoryID, productID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, options)
}

func GetOption(c echo.Context) error {
	id := c.Param("option_id")
	option, err := models.GetOption(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, id+" bulunamadı.")
	}
	return c.JSON(http.StatusOK, option)
}

// func UpdateOption(c echo.Context) error {
// 	id := c.Param("option_id")
// 	option, getErr := models.WithChoices(id)
// 	if getErr != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, getErr.Error())
// 	}
// 	err := c.Bind(&option)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
//
// 	_, updateErr := option.Update()
// 	if updateErr != nil {
// 		return updateErr
// 	}
// 	return c.JSON(http.StatusOK, option)
// }

// func CreateOption(c echo.Context) error {
// 	var category models.Option
// 	err := c.Bind(&category)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	createdOption, err := category.Create() // models'de metot
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(http.StatusCreated, createdOption)
// }
//
// func DeleteOption(c echo.Context) error {
// 	id := c.Param("category_id")
// 	category, err := models.GetOption(id)
// 	if err != nil {
// 		return c.String(http.StatusNotFound, id+"bulunamadı.")
// 	}
// 	category.Delete()
// 	return c.NoContent(http.StatusNoContent)
// }
//
// func UpdateOption(c echo.Context) error {
// 	id := c.Param("category_id")
// 	category, getErr := models.GetOption(id)
// 	if getErr != nil {
// 		return echo.NewHTTPError(http.StatusNotFound, getErr.Error())
// 	}
// 	err := c.Bind(&category)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
//
// 	_, updateErr := category.Update()
// 	if updateErr != nil {
// 		return updateErr
// 	}
// 	return c.JSON(http.StatusOK, category)
// }
