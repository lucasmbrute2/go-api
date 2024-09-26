package entity

type UserEntity struct {
	ID int `json:"id" gorm:"primaryKey"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age int `json:"age" validate:"required,number"`
	IsAdmin bool `json:"isAdmin"`
}