package services

import (
	"fmt"

	"github.com/NinePTH/GO_MVC-S/src/models"
)

func GetAllPatients() ([]models.Patient, error) {
	fields := []string{"*"}
	results, err := SelectData("Patients", fields, false, "", nil)
	if err != nil {
		return nil, err
	}
	var patients []models.Patient
	for _, row := range results {
		id := int(row["id"].(int64))
		first_name := string(row["first_name"].(string))
		last_name := string(row["last_name"].(string))
		age := int(row["age"].(int64))
		disease := string(row["disease"].(string))
		medicine := string(row["medicine"].(string))
		allergies := string(row["allergies"].(string))

		patient := models.Patient{
			Id:        id,
			First_name:      first_name,
			Last_name:   last_name,
			Age:       age,
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


