package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/models"
	"github.com/NinePTH/GO_MVC-S/src/services"

	"github.com/labstack/echo/v4"
)

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

    var req models.AddPatient
    if err:= c.Bind(&req); err != nil || req.Patient_id == "" || req.First_name == "" || req.Last_name == "" || req.Age == 0 || req.Gender == "" || req.Date_of_birth == "" || req.Blood_type == "" || req.Email == "" || req.Address == "" || req.Phone_number == "" || req.Id_card_number == "" || req.Ongoing_treatment == "" {
        return c.JSON(http.StatusBadRequest, "Invalid request, all information must be provided")
    }

	patientMap := map[string]interface{}{
		"patient_id": req.Patient_id,
		"first_name": req.First_name,
		"last_name": req.Last_name,
		"age": req.Age,
		"gender": req.Gender,
		"date_of_birth": req.Date_of_birth,
		"blood_type": req.Blood_type,
		"email": req.Email,
		"health_insurance": req.Health_insurance,
		"address": req.Address,
		"phone_number": req.Phone_number,
		"id_card_number": req.Id_card_number,
		"ongoing_treatment": req.Ongoing_treatment,
	}

    rowsAffected, err := services.AddPatient(patientMap)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, err.Error())
    }

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, "No rows affected")
	}

    return c.JSON(http.StatusOK, "Patient added successfully")
}