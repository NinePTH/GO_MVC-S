package patients
// import (
// 	"time"
// )
type AddPatientHistory struct{
Patient_id string `json:"patient_id"`
Detail string `json:"detail"`
Time string `json:"time"`
Date string `json:"date"`
}

// CREATE TABLE Medical_history (
// 	medical_history_id SERIAL PRIMARY KEY,
// 	patient_id VARCHAR(4) NOT NULL,
// 	detail TEXT NOT NULL,
// 	time TIME NOT NULL,
// 	date date NOT NULL,
// 	FOREIGN KEY (patient_id) REFERENCES Patient(patient_id)
// );