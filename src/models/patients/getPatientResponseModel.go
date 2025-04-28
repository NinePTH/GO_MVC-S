package patients

type GetPatientResponse struct {
	PatientGeneralInfo GeneralPatientInformation `json:"patient"`
	PatientMedicalHistory []MedicalHistory `json:"medical_history"`
	PatientChronicDisease []ChronicDiseaseName      `json:"patient_chronic_disease"`
	PatientDrugAllergy    []DrugAllergyName         `json:"patient_drug_allergy"`
}