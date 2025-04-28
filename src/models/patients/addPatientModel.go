// hospital patient database
package patients

type AddPatientRequest struct {
	Patient              GeneralPatientInformation            `json:"patient"`
	PatientChronicDisease []ChronicDiseaseName `json:"patient_chronic_disease"`
	PatientDrugAllergy    []DrugAllergyName    `json:"patient_drug_allergy"`
}