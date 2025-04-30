// hospital patient database
package patients

type SearchPatient struct {
	Patient_id              string            `json:"patient_id"`
	First_name             string            `json:"first_name"`
	Last_name              string            `json:"last_name"`	
}