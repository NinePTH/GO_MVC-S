package services

import (
	"fmt"

	"github.com/NinePTH/GO_MVC-S/src/models"
)
func UpdatePatient(id string,data map[string]interface{}) (int64, error){
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
func GetPatient(id string) (*models.Patient, error) {
	table := "Patient"
	fields := []string{"*"}

	result, err := SelectData(table, fields, true, "patient_id = $1", []interface{}{id})

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
	blood_type := string(result[0]["blood_type"].([]uint8))
	email := result[0]["email"].(string)
	health_insurance := result[0]["health_insurance"].(bool)
	address := result[0]["address"].(string)
	phone_number := result[0]["phone_number"].(string)
	id_card_number := result[0]["id_card_number"].(string)
	ongoing_treatment := result[0]["ongoing_treatment"].(string)

	var patient = models.Patient{
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
	}

	return &patient, nil
}

func GetAllPatients() ([]models.Patient, error) {
	fields := []string{"*"}
	results, err := SelectData("Patient", fields, false, "", nil)
	if err != nil {
		return nil, err
	}
	var patients []models.Patient
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

		patient := models.Patient{
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
		}
		patients = append(patients, patient)
		fmt.Println(patient)
	}
	fmt.Println(patients)
	return patients, nil
}
