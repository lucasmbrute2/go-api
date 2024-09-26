package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasmbrute2/go-api/internal/modules/user/dto"
	"github.com/lucasmbrute2/go-api/internal/modules/user/entity"
	"gorm.io/gorm"
)

type UserController struct {
	Db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Db: db,
	}
}

func (u *UserController) Handle(c echo.Context) error {
	var user dto.User

	err := c.Bind(&user); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err = c.Validate(user); err != nil {
		return err
	}
	
	u.Db.Create(&user)
	
	userDomain := entity.UserEntity{
		ID: user.ID,
		Email: user.Email,
		IsAdmin: false,
	}

	return c.JSON(http.StatusCreated, userDomain)
}
