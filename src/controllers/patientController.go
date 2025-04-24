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

func UpdatePatient(c echo.Context) error {
	patient_id := c.QueryParam("patient_id")
	if patient_id == "" {
		return c.JSON(http.StatusBadRequest, "Missing user ID")
	}

	data := map[string]interface{}{}

	// ใช้ helper function เพื่อเช็คและเพิ่มค่าเข้า map เฉพาะที่ไม่ว่าง (ทำให้สามารถ update แค่บางค่าได้)
	addIfNotEmpty := func(key, value string) {
		if value != "" {
			data[key] = value
		}
	}

	addIfNotEmpty("first_name", c.QueryParam("first_name"))
	addIfNotEmpty("last_name", c.QueryParam("last_name"))
	addIfNotEmpty("age", c.QueryParam("age"))
	addIfNotEmpty("date_of_birth", c.QueryParam("date_of_birth"))
	addIfNotEmpty("gender", c.QueryParam("gender"))
	addIfNotEmpty("blood_type", c.QueryParam("blood_type"))
	addIfNotEmpty("email", c.QueryParam("email"))
	addIfNotEmpty("address", c.QueryParam("address"))
	addIfNotEmpty("phone_number", c.QueryParam("phone_number"))
	addIfNotEmpty("id_card_number", c.QueryParam("id_card_number"))
	addIfNotEmpty("ongoing_treatment", c.QueryParam("ongoing_treatment"))

	// จัดการ health_insurance แยก เพราะเป็น bool
	health_insurance := c.QueryParam("health_insurance")
	if health_insurance != "" {
		if health_insurance == "true" {
			data["health_insurance"] = true
		} else {
			data["health_insurance"] = false
		}
	}

	// ถ้าไม่มี field อะไรเลยใน data ให้คืนว่าไม่มีอะไรอัปเดต
	if len(data) == 0 {
		return c.JSON(http.StatusBadRequest, "No data to update")
	}

	rowsAffected, err := services.UpdatePatient(patient_id, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, "No rows affected")
	}

	return c.JSON(http.StatusOK, "Patient information updated successfully")
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
	if p.Patient_id == "" || p.First_name == "" || p.Last_name == "" || p.Age == 0 || p.Gender == "" || p.Date_of_birth == "" || p.Blood_type == "" || p.Email == "" || p.Address == "" || p.Phone_number == "" || p.Id_card_number == "" || p.Ongoing_treatment == "" || p.Unhealthy_habits =="" {
		return c.JSON(http.StatusBadRequest, "All patient fields must be provided")
	}

	err := services.AddPatient(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Patient added successfully")
}
