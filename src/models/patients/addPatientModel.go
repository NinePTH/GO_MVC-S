// hospital patient database
package patients

type AddPatientRequest struct {
	Patient              AddPatient             `json:"patient"`
	PatientChronicDisease []ChronicDiseaseInput `json:"Patient_chronic_disease"`
	PatientDrugAllergy    []DrugAllergyInput    `json:"Patient_drug_allergy"`
}

type AddPatient struct {
	Patient_id         string `json:"patient_id"`
	First_name         string `json:"first_name"`
	Last_name          string `json:"last_name"`
	Age                int    `json:"age"`
	Gender             string `json:"gender"`
	Date_of_birth      string `json:"date_of_birth"`
	Blood_type         string `json:"blood_type"`
	Email              string `json:"email"`
	Health_insurance   bool   `json:"health_insurance"`
	Address            string `json:"address"`
	Phone_number       string `json:"phone_number"`
	Id_card_number     string `json:"id_card_number"`
	Ongoing_treatment  string `json:"ongoing_treatment"`
	Unhealthy_habits   string `json:"unhealthy_habits"`
}

type ChronicDiseaseInput struct {
	DiseaseID string `json:"disease_id"`
}

type DrugAllergyInput struct {
	DrugID string `json:"drug_id"`
}
