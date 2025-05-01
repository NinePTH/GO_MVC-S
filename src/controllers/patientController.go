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

func SearchPatient(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}

	body, _ := io.ReadAll(c.Request().Body)
	fmt.Println("Raw Request Body:", string(body))
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body)) // reset body for Bind()

	var req patients.SearchPatient
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	patients, err := services.GetPatientSearch(req.Patient_id, req.First_name, req.Last_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, patients)
}

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

	// ดักว่าเป็น string มั้ย
	validateString := func(fieldName string, value interface{}) error {
		if _, ok := value.(string); !ok {
			return fmt.Errorf("%s must be a string", fieldName)
		}
		return nil
	}

	// เช็คใน patient fields ว่าทุกค่าเป็น string หรือไม่
	if err := validateString("patient.patient_id", req.Patient.Patient_id); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.first_name", req.Patient.First_name); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.last_name", req.Patient.Last_name); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.gender", req.Patient.Gender); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.date_of_birth", req.Patient.Date_of_birth); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.blood_type", req.Patient.Blood_type); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.email", req.Patient.Email); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.health_insurance", req.Patient.Health_insurance); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.address", req.Patient.Address); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.phone_number", req.Patient.Phone_number); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.id_card_number", req.Patient.Id_card_number); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.ongoing_treatment", req.Patient.Ongoing_treatment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validateString("patient.unhealthy_habits", req.Patient.Unhealthy_habits); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// เช็คใน patient_chronic_disease ว่าทุกค่าเป็น string หรือไม่
	for i, chronic := range req.PatientChronicDisease {
		if err := validateString(fmt.Sprintf("patient_chronic_disease[%d].disease_id", i), chronic.DiseaseID); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	// เช็คใน patient_drug_allergy ว่าทุกค่าเป็น string หรือไม่
	for i, allergy := range req.PatientDrugAllergy {
		if err := validateString(fmt.Sprintf("patient_drug_allergy[%d].drug_id", i), allergy.DrugID); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	// Age must not be negative
	if req.Patient.Age < 0 {
		return c.JSON(http.StatusBadRequest, "Invalid Age Value")
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

	//Age must not be negative
	if req.Patient.Age < 0 {
		return c.JSON(http.StatusBadRequest, "Invalid Age Value")
	}

	err := services.AddPatient(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Patient added successfully")
}
