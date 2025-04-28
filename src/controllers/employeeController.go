package controllers

import (
	"fmt"
	"net/http"

	"github.com/NinePTH/GO_MVC-S/src/models"
	"github.com/NinePTH/GO_MVC-S/src/services"

	"github.com/labstack/echo/v4"
)

func UpdateEmployee(c echo.Context) error {
	employee_id := c.QueryParam("employee_id")
	if employee_id == "" {
		return c.JSON(http.StatusBadRequest, "Missing Employee ID")
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
	addIfNotEmpty("position_id", c.QueryParam("position_id"))
	addIfNotEmpty("phone_number", c.QueryParam("phone_number"))
	addIfNotEmpty("department_id", c.QueryParam("department_id"))
	addIfNotEmpty("salary", c.QueryParam("salary"))
	addIfNotEmpty("email", c.QueryParam("email"))
	addIfNotEmpty("hire_date", c.QueryParam("hire_date"))
	addIfNotEmpty("resignation_date", c.QueryParam("resignation_date"))
	addIfNotEmpty("work_status", c.QueryParam("work_status"))

	// ถ้าไม่มี field อะไรเลยใน data ให้คืนว่าไม่มีอะไรอัปเดต
	if len(data) == 0 {
		return c.JSON(http.StatusBadRequest, "No data to update")
	}

	rowsAffected, err := services.UpdateEmployee(employee_id, data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusOK, "No rows affected")
	}

	return c.JSON(http.StatusOK, "Employee information updated successfully")
}

func AddEmployee(c echo.Context) error {
	// ตรวจสอบ Content-Type
	if c.Request().Header.Get("Content-Type") != "application/json" {
		return c.JSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}

	var req models.Employee
	// Bind ข้อมูล JSON เข้าสู่ struct
	if err := c.Bind(&req); err != nil {
		fmt.Println("Error binding request:", err) // เพิ่ม log เพื่อดูข้อผิดพลาด
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// เช็คว่าทุก field ที่จำเป็นต้องมีค่าหมด
	if req.Employee_id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Employee ID is required"})
	}
	if req.First_name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "First name is required"})
	}
	if req.Last_name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Last name is required"})
	}
	if req.Position_name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Position name is required"})
	}
	if req.Phone_number == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Phone number is required"})
	}
	if req.Department_name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Department name is required"})
	}
	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email is required"})
	}
	if req.Hire_date == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Hire date is required"})
	}
	if req.Work_status == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Work status is required"})
	}
	if req.Salary == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Salary is required"})
	}

	// เตรียมข้อมูลที่จะ Insert
	data := map[string]interface{}{
		"employee_id":      req.Employee_id,
		"first_name":       req.First_name,
		"last_name":        req.Last_name,
		"position_name":    req.Position_name, // ใช้ position_name ใน JSON
		"phone_number":     req.Phone_number,
		"department_name":  req.Department_name, // ใช้ department_name ใน JSON
		"salary":           req.Salary,
		"email":            req.Email,
		"hire_date":        req.Hire_date,
		"resignation_date": nil, // default เป็น nil ก่อน
		"work_status":      req.Work_status,
	}

	// ถ้า resignation_date มีค่า (ไม่ว่าง) ให้ใส่
	if req.Resignation_date != "" {
		data["resignation_date"] = req.Resignation_date
	}

	// Insert ข้อมูลไปยังฐานข้อมูล
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
