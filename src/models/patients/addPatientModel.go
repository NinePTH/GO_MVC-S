// hospital patient database
package patients

type AddPatientRequest struct {
	Patient              GeneralPatientInformation            `json:"patient"`
	PatientChronicDisease []ChronicDiseaseName `json:"Patient_chronic_disease"`
	PatientDrugAllergy    []DrugAllergyName    `json:"Patient_drug_allergy"`
}