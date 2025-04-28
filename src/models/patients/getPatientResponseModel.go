package patients

type GetPatientResponse struct {
	PatientGeneralInfo GeneralPatientInformation `json:"patient"`
	PatientMedicalHistory []MedicalHistory `json:"medical_history"`
	PatientChronicDisease []ChronicDiseaseName      `json:"Patient_chronic_disease"`
	PatientDrugAllergy    []DrugAllergyName         `json:"Patient_drug_allergy"`
}