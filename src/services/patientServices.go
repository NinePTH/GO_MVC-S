package services

import (
	"fmt"
	"learn-Go/src/models"
)

func GetAllPatients() ([]models.Patient, error) {
	fields := []string{"id", "name", "surname", "age", "disease", "medicine", "allergies"}
	results, err := SelectData("patient", fields, false, "", nil)
	if err != nil {
		return nil, err
	}
	var patients []models.Patient
	for _, row := range results {
		id := int(row["ID"].(int))
		name := string(row["Name"].(string))
		surname := string(row["Surname"].(string))
		age := int(row["Age"].(int))
		disease := string(row["Disease"].(string))
		medicine := string(row["Medicine"].(string))
		allergies := string(row["Allergies"].(string))

		patient := models.Patient{
			ID:        int(id),
			Name:      name,
			Surname:   surname,
			Age:       int(age),
			Disease:   disease,
			Medicine:  medicine,
			Allergies: allergies,
		}
		patients = append(patients, patient)
		fmt.Println(patient)
	}
	fmt.Println(patients)
	return patients, nil
}


