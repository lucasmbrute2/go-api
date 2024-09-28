package dto

type Login struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}