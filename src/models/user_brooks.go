// hospital patient database
package models

type Patient struct {
	Id        int    `json:"id"`
	First_name      string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Age       int    `json:"age"`
	Disease   string `json:"disease"`
	Medicine  string `json:"medicine"`
	Allergies string `json:"allergies"`
}
