package patients
type AddPatientRequest struct {
	Patient                AddPatient           `json:"patient"`
	PatientChronicDisease  []ChronicDiseaseInput `json:"Patient_chronic_disease"`
	PatientDrugAllergy     []DrugAllergyInput    `json:"Patient_drug_allergy"`
}
type ChronicDiseaseInput struct {
	DiseaseID string `json:"disease_id"`
}
type DrugAllergyInput struct {
	DrugID string `json:"drug_id"`
}