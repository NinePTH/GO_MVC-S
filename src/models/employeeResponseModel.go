package models

import "time"

type EmployeeDetail struct {
	Employee_id      string    `json:"employee_id"`
	First_name       string    `json:"first_name"`
	Last_name        string    `json:"last_name"`
	Position_id      string    `json:"position_id"`
	Position_name    string    `json:"position_name"`
	Phone_number     string    `json:"phone_number"`
	Department_id    string    `json:"department_id"`
	Department_name  string    `json:"department_name"`
	Salary           float64   `json:"salary"`
	Email            string    `json:"email"`
	Hire_date        time.Time `json:"hire_date"`
	Resignation_date time.Time `json:"resignation_date"`
	Work_status      string    `json:"work_status"`
}
