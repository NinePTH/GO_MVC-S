package controllers

import (
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/services"

	"github.com/labstack/echo/v4"
)

func GetAllPatients(c echo.Context) error {
	patient, err := services.GetAllPatients()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, patient)
}
