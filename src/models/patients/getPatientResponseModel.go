package patients

type GetPatientResponse struct {
	PatientGeneralInfo GeneralPatientInformation `json:"patient"`
	PatientAppointment []PatientAppointment `json:"patient_appointment"`
	PatientMedicalHistory []MedicalHistory `json:"medical_history"`
	PatientChronicDisease []ChronicDiseaseName      `json:"patient_chronic_disease"`
	PatientDrugAllergy    []DrugAllergyName         `json:"patient_drug_allergy"`
}