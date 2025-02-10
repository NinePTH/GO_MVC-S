package models

type userAuthen struct{
	Id   string    `json:"id"`
	Name string `json:"name"`
	Password string `json:"password,omitempty"` // Omitting password in JSON response for security
}