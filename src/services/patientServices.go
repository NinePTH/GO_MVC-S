package services

import (
	"fmt"
	"time"

	"github.com/NinePTH/GO_MVC-S/src/models/patients"
)

func UpdatePatient(id string, data map[string]interface{}) (int64, error) {
	table := "Patient"
	condition := "patient_id = $1"
	conditionValues := []interface{}{id}

	// Call UpdateData with correct parameters
	rowsAffected, err := UpdateData(table, data, condition, conditionValues)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func AddPatient(data map[string]interface{}) (int64, error) {
	table := "Patient"
	rowsAffected, err := InsertData(table, data)
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
func GetPatient(id string) (*patients.GetPatientResponse, error) {
	table := "Patient"
	fields := []string{"*"}

	result, err := SelectData(table, fields, true, "patient_id = $1", []interface{}{id}, false, "", "")

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Patient not found")
	}

	patient_id := result[0]["patient_id"].(string)
	first_name := result[0]["first_name"].(string)
	last_name := result[0]["last_name"].(string)
	age := int(result[0]["age"].(int64))
	gender := string(result[0]["gender"].([]uint8))
	date_of_birth := string(result[0]["date_of_birth"].(time.Time).Format("02-01-2006"))
	blood_type := string(result[0]["blood_type"].([]uint8))
	email := result[0]["email"].(string)
	health_insurance := result[0]["health_insurance"].(bool)
	address := result[0]["address"].(string)
	phone_number := result[0]["phone_number"].(string)
	id_card_number := result[0]["id_card_number"].(string)
	ongoing_treatment := result[0]["ongoing_treatment"].(string)
	unhealthy_habits := result[0]["unhealthy_habits"].(string)

	var patient = patients.GeneralPatientInformation{
		Patient_id:        patient_id,
		First_name:        first_name,
		Last_name:         last_name,
		Age:               age,
		Gender:            gender,
		Date_of_birth:     date_of_birth,
		Blood_type:        blood_type,
		Email:             email,
		Health_insurance:  health_insurance,
		Address:           address,
		Phone_number:      phone_number,
		Id_card_number:    id_card_number,
		Ongoing_treatment: ongoing_treatment,
		Unhealthy_habits: unhealthy_habits,
		
	}

	table = "Medical_history"
	fields = []string{"*"}

	result, err = SelectData(table, fields, true, "patient_id = $1", []interface{}{id},false,"","")

	if err != nil {
		return nil, err
	}

	var medical_history []patients.MedicalHistory
	for _, row := range result {
		details := row["detail"].(string)
		date := string(row["date"].(time.Time).Format("02-01-2006"))
		time := string(row["time"].(time.Time).Format("02:01:00"))

		medical_history = append(medical_history, patients.MedicalHistory{
			Details: details,
			Date:    date,
			Time:    time,
		})
	}

	var response = patients.GetPatientResponse{
		PatientGeneralInfo: patient,
		PatientMedicalHistory: medical_history,
	}

	return &response, nil
}

func GetAllPatients() ([]patients.GeneralPatientInformation, error) {
	fields := []string{"*"}
	results, err := SelectData("Patient", fields, false, "", nil, false, "", "")
	if err != nil {
		return nil, err
	}
	var patientList []patients.GeneralPatientInformation
	for _, row := range results {
		patient_id := row["patient_id"].(string)
		first_name := row["first_name"].(string)
		last_name := row["last_name"].(string)
		age := int(row["age"].(int64))
		gender := string(row["gender"].([]uint8))
		blood_type := string(row["blood_type"].([]uint8))
		email := row["email"].(string)
		health_insurance := row["health_insurance"].(bool)
		address := row["address"].(string)
		phone_number := row["phone_number"].(string)
		id_card_number := row["id_card_number"].(string)
		ongoing_treatment := row["ongoing_treatment"].(string)
		unhealthy_habits := row["unhealthy_habits"].(string)

		patient := patients.GeneralPatientInformation{
			Patient_id:        patient_id,
			First_name:        first_name,
			Last_name:         last_name,
			Age:               age,
			Gender:            gender,
			Blood_type:        blood_type,
			Email:             email,
			Health_insurance:  health_insurance,
			Address:           address,
			Phone_number:      phone_number,
			Id_card_number:    id_card_number,
			Ongoing_treatment: ongoing_treatment,
			Unhealthy_habits: unhealthy_habits,
		}
		patientList = append(patientList, patient)
		fmt.Println(patient)
	}
	fmt.Println(patientList)
	return patientList, nil
}