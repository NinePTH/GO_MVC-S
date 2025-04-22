package patients

type GetPatientResponse struct {
	PatientGeneralInfo GeneralPatientInformation `json:"patient"`
	PatientMedicalHistory []MedicalHistory `json:"medical_history"`
}