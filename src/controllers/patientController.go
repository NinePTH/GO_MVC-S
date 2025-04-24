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

	patient_id := req.Patient.Patient_id // ใช้ patient_id จาก body ที่ส่งมา
	if patient_id == "" {
		return c.JSON(http.StatusBadRequest, "Missing patient_id")
	}

	data := map[string]interface{}{}
	addIfNotEmpty := func(key, value string) {
		if value != "" {
			data[key] = value
		}
	}

	// เพิ่มข้อมูลที่ไม่ว่างใน data
	addIfNotEmpty("first_name", req.Patient.First_name)
	addIfNotEmpty("last_name", req.Patient.Last_name)
	addIfNotEmpty("age", fmt.Sprintf("%v", req.Patient.Age)) // แปลง age เป็น string (รับเป็น null)
	addIfNotEmpty("date_of_birth", req.Patient.Date_of_birth)
	addIfNotEmpty("gender", req.Patient.Gender)
	addIfNotEmpty("blood_type", req.Patient.Blood_type)
	addIfNotEmpty("email", req.Patient.Email)
	addIfNotEmpty("address", req.Patient.Address)
	addIfNotEmpty("phone_number", req.Patient.Phone_number)
	addIfNotEmpty("id_card_number", req.Patient.Id_card_number)
	addIfNotEmpty("ongoing_treatment", req.Patient.Ongoing_treatment)
	addIfNotEmpty("unhealthy_habits", req.Patient.Unhealthy_habits)

	// Handle health_insurance (boolean)
	data["health_insurance"] = req.Patient.Health_insurance
	

	// Update patient info
	if len(data) > 0 {
		rowsAffected, err := services.UpdatePatient(patient_id, data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		fmt.Printf("Patient info updated: %d rows affected\n", rowsAffected)
	}

	// ============ Chronic Diseases ============ (ถ้าส่งเปล่า->ไม่เกิดอะไรขึ้น เก็บค่าเดิม) (ถ้าส่งค่าจะ rewrite ค่าเก่า)
	// ถ้ามีการใส่ค่าเข้า arrray PatientChronicDisease
	if len(req.PatientChronicDisease) > 0 {
		// ลบ ค่าใน table patient_chronic_disease เก่าที่มีอยู่
		table := "patient_chronic_disease"
		if err := services.DeleteByPatientID(table, patient_id); err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to delete chronic diseases")
		}

		// เพิ่ม chronic diseases ใหม่
		for _, chronic := range req.PatientChronicDisease {
			chronicMap := map[string]interface{}{
				"patient_id": patient_id,
				"disease_id": chronic.DiseaseID,
			}
			_, err := services.InsertData("patient_chronic_disease", chronicMap)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Insert chronic disease failed: %v", err))
			}
		}
	}

	// ============ Drug Allergies ============ (ถ้าส่งเปล่า->ไม่เกิดอะไรขึ้น เก็บค่าเดิม) (ถ้าส่งค่าจะ rewrite ค่าเก่า)
	// ถ้ามีการใส่ค่าเข้า arrray PatientChronicDisease
	if len(req.PatientDrugAllergy) > 0 {
		// ลบ ค่าใน table patient_drug_allergy drug allergies เก่าที่มีอยู่
		table := "patient_drug_allergy"
		if err := services.DeleteByPatientID(table, patient_id); err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to delete drug allergies")
		}

		// เพิ่ม drug allergies ใหม่
		for _, allergy := range req.PatientDrugAllergy {
			allergyMap := map[string]interface{}{
				"patient_id": patient_id,
				"drug_id":    allergy.DrugID,
			}
			_, err := services.InsertData("patient_drug_allergy", allergyMap)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Insert drug allergy failed: %v", err))
			}
		}
	}

	return c.JSON(http.StatusOK, "Patient updated successfully")
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
