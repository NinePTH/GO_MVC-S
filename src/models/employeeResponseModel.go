package models

//import "time"

type EmployeeResponse struct {
	Employee_id      string    `json:"employee_id"`
	First_name       string    `json:"first_name"`
	Last_name        string    `json:"last_name"`
	Position_name    string    `json:"position_name"`
	Department_name    string    `json:"department_name"`
	Phone_number     string    `json:"phone_number"`
	Salary           float64   `json:"salary"`
	Email            string    `json:"email"`
	Hire_date        string `json:"hire_date"`
	Resignation_date string `json:"resignation_date"`
	Work_status      string    `json:"work_status"`
}
