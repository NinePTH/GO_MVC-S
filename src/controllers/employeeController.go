package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/models"
	//"github.com/NinePTH/GO_MVC-S/src/models/patients"
	"github.com/NinePTH/GO_MVC-S/src/services"

	"github.com/labstack/echo/v4"
)

func SearchEmployee(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}

	body, _ := io.ReadAll(c.Request().Body)
	fmt.Println("Raw Request Body:", string(body))
	c.Request().Body = io.NopCloser(bytes.NewBuffer(body)) // reset body for Bind()

	var req models.SearchEmployee
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	patients, err := services.GetEmployeeSearch(req.Employee_id, req.First_name, req.Last_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, patients)
}
func UpdateEmployee(c echo.Context) error {
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}

	var raw map[string]interface{}
	if err := c.Bind(&raw); err != nil {
		fmt.Println("Error binding request:", err)
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// ตรวจสอบว่า employee_id มี และเป็น string
	employeeID, ok := raw["employee_id"].(string)
	if !ok || employeeID == "" {
		return c.JSON(http.StatusBadRequest, "Missing or invalid employee_id")
	}

	data := map[string]interface{}{}

	//ดัก undefined ทำให้รับได้แค่ string กับ float
	addIfValidString := func(key string) {
		if val, ok := raw[key]; ok {
			str, ok := val.(string)
			if !ok {
				c.JSON(http.StatusBadRequest, fmt.Sprintf("%s must be a string", key))
				return
			}
			if str != "" {
				data[key] = str
			}
		}
	}
	addIfValidFloat := func(key string) {
		if val, ok := raw[key]; ok {
			num, ok := val.(float64)
			if !ok {
				c.JSON(http.StatusBadRequest, fmt.Sprintf("%s must be a number", key))
				return
			}
			if num != 0 {
				data[key] = num
			}
		}
	}

	addIfValidString("first_name")
	addIfValidString("last_name")
	addIfValidString("position_id")
	addIfValidString("phone_number")
	addIfValidString("email")
	addIfValidString("hire_date")
	addIfValidString("work_status")
	addIfValidString("resignation_date")
	addIfValidFloat("salary")

	if len(data) == 0 {
		return c.JSON(http.StatusBadRequest, "No valid data to update")
	}

	rowsAffected, err := services.UpdateEmployee(employeeID, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, map[string]string{"message": "No rows affected"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Employee information updated successfully"})
}

// func UpdateEmployee(c echo.Context) error {
// 	// ตรวจสอบ Content-Type
// 	if c.Request().Header.Get("Content-Type") != "application/json" {
// 		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
// 	}

// 	var req models.EmployeeInsert
// 	if err := c.Bind(&req); err != nil {
// 		fmt.Println("Error binding request:", err)
// 		return c.JSON(http.StatusBadRequest, "Invalid request body")
// 	}

// 	employee_id := req.Employee_id
// 	if employee_id == "" {
// 		return c.JSON(http.StatusBadRequest, "Missing Employee ID")
// 	}
// 	data := map[string]interface{}{}

// 	// ใช้ฟังก์ชัน addIfNotEmpty เพื่อเพิ่มเฉพาะ field ที่มีค่า
// 	addIfNotEmpty := func(key string, value interface{}) {
// 		switch v := value.(type) {
// 		case string:
// 			if v != "" {
// 				data[key] = v
// 			}
// 		case float64:
// 			if v != 0 {
// 				data[key] = v
// 			}
// 		}
// 	}

// 	addIfNotEmpty("first_name", req.First_name)
// 	addIfNotEmpty("last_name", req.Last_name)
// 	addIfNotEmpty("position_id", req.Position_id)
// 	addIfNotEmpty("phone_number", req.Phone_number)
// 	addIfNotEmpty("email", req.Email)
// 	addIfNotEmpty("hire_date", req.Hire_date)
// 	addIfNotEmpty("work_status", req.Work_status)
// 	addIfNotEmpty("salary", req.Salary)
// 	addIfNotEmpty("resignation_date", req.Resignation_date)

// 	if len(data) == 0 {
// 		return c.JSON(http.StatusBadRequest, "No data to update")
// 	}


// 	rowsAffected, err := services.UpdateEmployee(employee_id, data)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 	}

// 	if rowsAffected == 0 {
// 		return c.JSON(http.StatusOK, map[string]string{"message": "No rows affected"})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{"message": "Employee information updated successfully"})
// }

func AddEmployee(c echo.Context) error { // แยก model ตอนส่งกับรับกลับ ส่ง id รับ name
	// ตรวจสอบ Content-Type
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}

	var req models.EmployeeInsert
	if err := c.Bind(&req); err != nil {
		fmt.Println("Error binding request:", err)
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// ตรวจสอบ fields ที่จำเป็น
	if req.Employee_id == "" || req.First_name == "" || req.Last_name == "" ||
		req.Position_id == "" || req.Phone_number == "" ||
		req.Email == "" || req.Hire_date == "" || req.Work_status == "" || req.Salary == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
	}

	// เตรียมข้อมูล insert
	data := map[string]interface{}{
		"employee_id":      req.Employee_id,
		"first_name":       req.First_name,
		"last_name":        req.Last_name,
		"position_id":      req.Position_id,
		"phone_number":     req.Phone_number,
		"salary":           req.Salary,
		"email":            req.Email,
		"hire_date":        req.Hire_date,
		"resignation_date": nil,
		"work_status":      req.Work_status,
	}

	if req.Resignation_date != "" {
		data["resignation_date"] = req.Resignation_date
	}

	rowsAffected, err := services.AddEmployee(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, map[string]string{"message": "No rows affected"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Employee added successfully"})
}

func GetAllEmployee(c echo.Context) error {
	employee, err := services.GetAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, employee)
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
