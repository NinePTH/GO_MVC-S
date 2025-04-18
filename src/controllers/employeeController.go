package controllers

import (
	"net/http"

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
	employee_id := c.QueryParam("employee_id")
	if employee_id == "" {
		return c.JSON(http.StatusBadRequest, "Missing employee ID")
	}
	
	first_name := c.QueryParam("first_name")
	last_name := c.QueryParam("last_name")
	position_id := c.QueryParam("position_id")
	phone_number := c.QueryParam("phone_number")
	department_id := c.QueryParam("department_id")
	salary := c.QueryParam("salary")
	email := c.QueryParam("email")
	hire_date := c.QueryParam("hire_date")
	resignation_date := c.QueryParam("resignation_date")
	var resignationDate interface{}
	if resignation_date == "" {
		resignationDate = nil
	} else {
		resignationDate = resignation_date
	}
	work_status := c.QueryParam("work_status")

	//ดัก null ทุกช่อง เพราะเพิ่มประวัติพนักงานต้องกรอกข้อมูลให้ครบ
	if first_name == "" {
		return c.JSON(http.StatusBadRequest, "Missing First Name")
	}
	if last_name == "" {
		return c.JSON(http.StatusBadRequest, "Missing Last Name")
	}
	if position_id == "" {
		return c.JSON(http.StatusBadRequest, "Missing Position Id")
	}
	if phone_number == "" {
		return c.JSON(http.StatusBadRequest, "Missing Phone Number")
	}
	if department_id == "" {
		return c.JSON(http.StatusBadRequest, "Missing Department id")
	}
	if salary == "" {
		return c.JSON(http.StatusBadRequest, "Missing Salary")
	}
	if email == "" {
		return c.JSON(http.StatusBadRequest, "Missing Email")
	}
	if hire_date == "" {
		return c.JSON(http.StatusBadRequest, "Missing Hire Date")
	}
	//ไม่เช็ค resignation date
	if work_status == "" {
		return c.JSON(http.StatusBadRequest, "Missing Work Status")
	}
	data := map[string]interface{}{
		"employee_id":      employee_id,
		"first_name":       first_name,
		"last_name":        last_name,
		"position_id":      position_id,
		"phone_number":     phone_number,
		"department_id":    department_id,
		"salary":           salary,
		"email":            email,
		"hire_date":        hire_date,
		"resignation_date": resignationDate, // optional
		"work_status":      work_status,
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
	patient, err := services.GetAllEmployee()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, patient)
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
