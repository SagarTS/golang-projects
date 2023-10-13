package models

type UserRegistration struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Age int `json:"age"`
	Address string `json:"address"`
}

func (b *UserRegistration) TableName() string {
	return "user_registration"
}

