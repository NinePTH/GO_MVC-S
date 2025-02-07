// hospital patient database
package models

type Patient struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Disease   string `json:"disease"`
	Medicine  string `json:"medicine"`
	Allergies string `json:"allergies"`
}
