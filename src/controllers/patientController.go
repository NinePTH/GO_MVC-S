package controllers

import (
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/services"

	"github.com/labstack/echo/v4"
)

func UpdatePatient(c echo.Context) error {
	patient_id := c.QueryParam("patient_id")

	if patient_id == "" {
		return c.JSON(http.StatusBadRequest, "Missing user ID")
	}

	first_name := c.QueryParam("first_name")
	last_name := c.QueryParam("last_name")
	age := c.QueryParam("age")
	date_of_birth := c.QueryParam("date_of_birth")
	gender := c.QueryParam("gender")
	blood_type := c.QueryParam("blood_type")
	email := c.QueryParam("email")
	health_insurance := c.QueryParam("health_insurance")
	address := c.QueryParam("address")
	phone_number := c.QueryParam("phone_number")
	id_card_number := c.QueryParam("id_card_number")
	ongoing_treatment := c.QueryParam("ongoing_treatment")
	//ดัก null
	if first_name == "" {
		return c.JSON(http.StatusBadRequest, "Missing First Name")
	}
	if last_name == "" {
		return c.JSON(http.StatusBadRequest, "Missing Last Name")
	}
	if age == "" {
		return c.JSON(http.StatusBadRequest, "Missing Age")
	}
	if date_of_birth == "" {
		return c.JSON(http.StatusBadRequest, "Missing Date Of Birth")
	}
	if gender == "" {
		return c.JSON(http.StatusBadRequest, "Missing Gender")
	}
	if blood_type == "" {
		return c.JSON(http.StatusBadRequest, "Missing Blood Type")
	}
	if email == "" {
		return c.JSON(http.StatusBadRequest, "Missing Email")
	}
	if health_insurance == "" {
		return c.JSON(http.StatusBadRequest, "Missing Health Insurance")
	}
	if address == "" {
		return c.JSON(http.StatusBadRequest, "Missing Address")
	}
	if phone_number == "" {
		return c.JSON(http.StatusBadRequest, "Missing Phone Number")
	}
	if id_card_number == "" {
		return c.JSON(http.StatusBadRequest, "Missing Id Card Number")
	}
	if ongoing_treatment == "" {
		return c.JSON(http.StatusBadRequest, "Missing Ongoing Treatment")
	}
	// แปลง health_insurance จาก string -> bool
	insuranceBool := false
	if health_insurance == "true" {
		insuranceBool = true
	}

	data := map[string]interface{}{
		"patient_id":        patient_id,
		"first_name":        first_name,
		"last_name":         last_name,
		"age":               age,
		"date_of_birth":     date_of_birth,
		"gender":            gender,
		"blood_type":        blood_type,
		"email":             email,
		"health_insurance":  insuranceBool,
		"address":           address,
		"phone_number":      phone_number,
		"id_card_number":    id_card_number,
		"ongoing_treatment": ongoing_treatment,
	}
	rowsAffected, err := services.UpdatePatient(patient_id, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, "No rows affected")
	}

	return c.JSON(http.StatusOK, "User updated successfully")
}
func AddPatient(c echo.Context) error {
	patient_id := c.QueryParam("patient_id")
	first_name := c.QueryParam("first_name")
	last_name := c.QueryParam("last_name")
	age := c.QueryParam("age")
	date_of_birth := c.QueryParam("date_of_birth")
	gender := c.QueryParam("gender")
	blood_type := c.QueryParam("blood_type")
	email := c.QueryParam("email")
	health_insurance := c.QueryParam("health_insurance")
	address := c.QueryParam("address")
	phone_number := c.QueryParam("phone_number")
	id_card_number := c.QueryParam("id_card_number")
	ongoing_treatment := c.QueryParam("ongoing_treatment")

	// แปลง health_insurance จาก string -> bool
	insuranceBool := false
	if health_insurance == "true" {
		insuranceBool = true
	}

	data := map[string]interface{}{
		"patient_id":        patient_id,
		"first_name":        first_name,
		"last_name":         last_name,
		"age":               age,
		"date_of_birth":     date_of_birth,
		"gender":            gender,
		"blood_type":        blood_type,
		"email":             email,
		"health_insurance":  insuranceBool,
		"address":           address,
		"phone_number":      phone_number,
		"id_card_number":    id_card_number,
		"ongoing_treatment": ongoing_treatment,
	}

	rowsAffected, err := services.UpdatePatient(patient_id, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, map[string]string{"message": "No rows affected"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Patient added successfully"})
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
	employee, err := services.GetAllPatients()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, employee)
}
