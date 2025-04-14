package controllers

import (
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/services"

	"github.com/labstack/echo/v4"
)

func GetAllEmployee(c echo.Context) error {
	patient, err := services.GetAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, patient)
}
func GetEmployee(c echo.Context) error {
	id := c.Param("id")
	user, err := services.GetEmployee(id)
	if err != nil {
		if err.Error() == "employee not found" {
			return c.JSON(http.StatusNoContent, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}