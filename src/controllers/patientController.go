// limit
package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/models/patients"
	"github.com/NinePTH/GO_MVC-S/src/services"
	"github.com/labstack/echo/v4"
)

func AddPatientAppointment(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}
	// Log raw request body
	body, _ := io.ReadAll(c.Request().Body)
	fmt.Println("Raw Request Body:", string(body))
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body)) // Reset body for Bind()

	var req patients.AddPatientAppointment

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// check all fields of patient must be filled
	if req.Patient_id == "" || req.Topic == "" || req.Time == "" || req.Date == "" {
		return c.JSON(http.StatusBadRequest, "All patient fields must be provided")
	}

	err := services.AddPatientAppointment(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Patient appointment added successfully")
}
func AddPatientHistory(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}
	// Log raw request body
	body, _ := io.ReadAll(c.Request().Body)
	fmt.Println("Raw Request Body:", string(body))
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body)) // Reset body for Bind()

	var req patients.AddPatientHistory

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// check all fields of patient must be filled
	if req.Patient_id == "" || req.Detail == "" || req.Time == "" || req.Date == "" {
		return c.JSON(http.StatusBadRequest, "All patient fields must be provided")
	}

	err := services.AddPatientHistory(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Patient history added successfully")
}

func UpdatePatient(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}

	var req patients.AddPatientRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	rowsAffected, err := services.UpdatePatient(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Patient updated successfully",
		"rows_affected": rowsAffected,
	})
}

func GetPatient(c echo.Context) error {
	id := c.Param("id")
	user, err := services.GetPatient(id)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNoContent, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func GetAllPatients(c echo.Context) error {
	patient, err := services.GetAllPatients()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, patient)
}
func AddPatient(c echo.Context) error {

	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}
	// Log raw request body
	body, _ := io.ReadAll(c.Request().Body)
	fmt.Println("Raw Request Body:", string(body))
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body)) // Reset body for Bind()

	var req patients.AddPatientRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// check all fields of patient must be filled
	p := req.Patient
	if p.Patient_id == "" || p.First_name == "" || p.Last_name == "" || p.Age == 0 || p.Gender == "" || p.Date_of_birth == "" || p.Blood_type == "" || p.Email == "" || p.Address == "" || p.Phone_number == "" || p.Id_card_number == "" || p.Ongoing_treatment == "" || p.Unhealthy_habits == "" {
		return c.JSON(http.StatusBadRequest, "All patient fields must be provided")
	}

	err := services.AddPatient(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Patient added successfully")
}
