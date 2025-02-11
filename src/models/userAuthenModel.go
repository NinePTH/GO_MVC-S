package models

type UserAuthen struct{
	Id   int    `json:"id"`
	Username string `json:"name"`
	Password string `json:"password,omitempty"` // Omitting password in JSON response for security
}