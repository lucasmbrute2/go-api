package view

import "github.com/lucasmbrute2/go-api/internal/modules/user/entity"

type UserView struct {
	ID int `json:"id"`
	Email string `json:"email" validate:"required,email"`
	Name string `json:"name" validate:"required"`
	Age int `json:"age" validate:"required,number"`
	IsAdmin bool `json:"isAdmin"`
}

func NewUserView() *UserView{
	return &UserView{}
}

func (u *UserView) ToHTTP(user entity.User) UserView {
	return UserView{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Age: user.Age,
		IsAdmin: user.IsAdmin,
	}
}

func (u *UserView) UpdateToHTTP(user entity.UpdateUser) UserView {
	return UserView{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Age: user.Age,
	}
}
