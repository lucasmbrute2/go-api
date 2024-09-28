package dto

type User struct {
	ID int `json:"id" gorm:"primaryKey"`
	Email string `json:"email" validate:"required,email" gorm:"unique;not null"`
	Name string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required" gorm:"not null"`
	Age int `json:"age" validate:"required,number" gorm:"not null"`
	IsAdmin bool `json:"isAdmin"`
}

type UpdateUser struct {
	ID int `json:"id"`
	Email string `json:"email" validate:"omitempty,email"`
	Name string `json:"name" validate:"omitempty"`
	Password string `json:"password"`
	Age int `json:"age" validate:"number"`
}