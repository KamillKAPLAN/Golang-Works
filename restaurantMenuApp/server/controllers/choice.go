package controllers

import (
	"net/http"

	"github.com/ridvansumset/restaurant-menu-app/server/models"

	"github.com/labstack/echo"
)

func ListChoices(c echo.Context) error {
	categoryID := c.Param("category_id")
	productID := c.Param("product_id")
	optionID := c.Param("option_id")
	choices, err := models.ListChoices(categoryID, productID, optionID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, choices)
}

func GetChoice(c echo.Context) error {
	id := c.Param("choice_id")
	choice, err := models.GetChoice(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, id+" bulunamadı.")
	}
	return c.JSON(http.StatusOK, choice)
}
