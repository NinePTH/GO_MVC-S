package controllers

import (
	"learn-Go/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllPatients(c echo.Context) error {
	patient, err := services.GetAllPatients()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, patient)
}
