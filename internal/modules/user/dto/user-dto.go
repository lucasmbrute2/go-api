package dto

type 	User struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age int `json:"age" validate:"required,number"`
}