package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucasmbrute2/go-api/internal/infra/cipher"
	"github.com/lucasmbrute2/go-api/internal/infra/jwt"
	"github.com/lucasmbrute2/go-api/internal/modules/auth/dto"
	userDto "github.com/lucasmbrute2/go-api/internal/modules/user/dto"
	"gorm.io/gorm"
)

type Login struct {
	Token string `json:"accessToken"`
}

var BadRequestError = "bad request error"
var InvalidCredentialsError = "invalid credentials"

type AuthController struct {
	Db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		Db: db,
	}
}

func (a *AuthController) Login(c echo.Context) error {
	var payload dto.Login

	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, BadRequestError)
	}

	if err := c.Validate(payload); err != nil {
		return err
	}

	var user userDto.User
	err := a.Db.Where("email = ?", payload.Username).First(&user).Error

	if err != nil {
		return c.String(http.StatusBadRequest, InvalidCredentialsError)
	}

	salts := 6
	cipher := cipher.NewCipher(salts)

	if ok, err := cipher.Compare(payload.Password, user.Password); err != nil || !ok {
		return c.String(http.StatusBadRequest, InvalidCredentialsError)
	}

	s := jwt.JWT{
		Secret: "secret",
	}

	token, err := s.Generate(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, Login{
		Token: token,
	})
}