package models


type LoginCredentials struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}