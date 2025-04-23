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

    var req patients.AddPatient
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

// func AddPatient(c echo.Context) error {
// 	if c.Request().Header.Get("Content-Type") != "application/json" {
// 		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
// 	}

// 	body, _ := io.ReadAll(c.Request().Body)
// 	fmt.Println("Raw Request Body:", string(body))
// 	c.Request().Body = io.NopCloser(bytes.NewBuffer(body)) // Reset body

// 	var req patients.AddPatientRequest
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, "Invalid request body")
// 	}

// 	// ตรวจสอบความครบถ้วนของข้อมูล patient
// 	p := req.Patient
// 	if p.Patient_id == "" || p.First_name == "" || p.Last_name == "" || p.Age == 0 ||
// 		p.Gender == "" || p.Date_of_birth == "" || p.Blood_type == "" || p.Email == "" ||
// 		p.Address == "" || p.Phone_number == "" || p.Id_card_number == "" || p.Ongoing_treatment == "" {
// 		return c.JSON(http.StatusBadRequest, "All patient fields must be provided")
// 	}

// 	// สร้าง map สำหรับ Insert
// 	patientMap := map[string]interface{}{
// 		"patient_id":        p.Patient_id,
// 		"first_name":        p.First_name,
// 		"last_name":         p.Last_name,
// 		"age":               p.Age,
// 		"gender":            p.Gender,
// 		"date_of_birth":     p.Date_of_birth,
// 		"blood_type":        p.Blood_type,
// 		"email":             p.Email,
// 		"health_insurance":  p.Health_insurance,
// 		"address":           p.Address,
// 		"phone_number":      p.Phone_number,
// 		"id_card_number":    p.Id_card_number,
// 		"ongoing_treatment": p.Ongoing_treatment,
// 	}

// 	// ➤ ส่งข้อมูลไปยัง services.AddPatient
// 	rowsAffected, err := services.AddPatient(patientMap,"patient") //table name
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, err.Error())
// 	}

// 	// ➤ เพิ่ม chronic disease ให้ผู้ป่วย
// 	for _, cd := range req.PatientChronicDisease {
// 		_, err := services.AddPatient(p.Patient_id, cd.DiseaseID)
// 		if err != nil {
// 			fmt.Println("Error adding chronic disease:", err)
// 		}
// 	}

// 	// ➤ เพิ่ม drug allergy ให้ผู้ป่วย
// 	for _, da := range req.PatientDrugAllergy {
// 		_, err := services.AddPatient(p.Patient_id, da.DrugID)
// 		if err != nil {
// 			fmt.Println("Error adding drug allergy:", err)
// 		}
// 	}

// 	if rowsAffected == 0 {
// 		return c.JSON(http.StatusOK, "No rows affected")
// 	}

// 	return c.JSON(http.StatusOK, "Patient added successfully")
// }
