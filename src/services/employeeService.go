package services

import (
	"database/sql" // เพิ่มการ import
	"fmt"
	"time"

	"github.com/NinePTH/GO_MVC-S/src/models"
)

// แก้ไขการดึงข้อมูลใน GetAllEmployee
func GetAllEmployee() ([]models.EmployeeDetail, error) {
	fields := []string{
		"Employee.employee_id",
		"Employee.first_name",
		"Employee.last_name",
		"Employee.position_id",
		"Position.position_name",
		"Employee.phone_number",
		"Employee.department_id",
		"Department.department_name",
		"Employee.salary",
		"Employee.email",
		"Employee.hire_date",
		"Employee.resignation_date",
		"Employee.work_status",
	}

	results, err := SelectInnerJoin(
		"Employee",
		"Department ON Employee.department_id = Department.department_id INNER JOIN Position ON Employee.position_id = Position.position_id",
		"",
		fields,
		false,
		"",
		nil,
	)

	if err != nil {
		return nil, err
	}

	var employees []models.EmployeeDetail

	// ลูปผลลัพธ์
	for _, row := range results {
		// จัดการกับ resignation_date ที่อาจเป็น NULL
		var resignationDate sql.NullTime
		if row["resignation_date"] != nil {
			resignationDate = sql.NullTime{Time: row["resignation_date"].(time.Time), Valid: true}
		}

		// แปลง salary ให้เป็น float64
		salary := parseSalary(row["salary"])

		// แปลง work_status ให้เป็น string
		workStatus := string(row["work_status"].([]byte))


		employee := models.EmployeeDetail{
			Employee_id:      fmt.Sprintf("%v", row["employee_id"]),
			First_name:       fmt.Sprintf("%v", row["first_name"]),
			Last_name:        fmt.Sprintf("%v", row["last_name"]),
			Position_id:      fmt.Sprintf("%v", row["position_id"]),
			Position_name:    fmt.Sprintf("%v", row["position_name"]),
			Phone_number:     fmt.Sprintf("%v", row["phone_number"]),
			Department_id:    fmt.Sprintf("%v", row["department_id"]),
			Department_name:  fmt.Sprintf("%v", row["department_name"]),
			Salary:           salary, // salary ที่แปลงแล้ว
			Email:            fmt.Sprintf("%v", row["email"]),
			Hire_date:        row["hire_date"].(time.Time),
			// ตรวจสอบว่า resignationDate.Valid เป็น true หรือไม่
			Resignation_date: resignationDate.Time, // ใช้ Time ถ้า Valid เป็น true
			Work_status:      workStatus, // work_status ที่แปลงแล้ว
		}

		employees = append(employees, employee)
		fmt.Println(employee)
	}

	return employees, nil
}

// แปลง salary เป็น float64
func parseSalary(data interface{}) float64 {
	salaryBytes := data.([]byte)   // แปลงเป็น []byte ก่อน
	return float64(salaryBytes[0]) // สมมติว่าค่า salary อยู่ในช่องที่ 0 ของ []byte
}
