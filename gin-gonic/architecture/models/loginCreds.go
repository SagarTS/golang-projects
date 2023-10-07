package models

type LoginCredentials struct{
	Email string `json:"email"`
	Password string `json:"password"`
}