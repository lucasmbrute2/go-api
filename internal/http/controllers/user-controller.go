package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lucasmbrute2/go-api/internal/http/view"
	"github.com/lucasmbrute2/go-api/internal/infra/cipher"
	"github.com/lucasmbrute2/go-api/internal/modules/user/dto"
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

func (u *UserController) CreateUser(c echo.Context) error {
	var user dto.User

	err := c.Bind(&user); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	fmt.Println(user.Email)
	
	var has dto.User
	err = u.Db.Where("email = ?", user.Email).First(&has).Error

	if !errors.Is(err, gorm.ErrRecordNotFound){
		return c.String(http.StatusBadRequest, "email already in use")
	}

	if err = c.Validate(user); err != nil {
		return err
	}
	
	salts := 6
	cipher := cipher.NewCipher(salts)
	cipherPassword, err :=  cipher.Encrypt(user.Password)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	user.Password = cipherPassword

	u.Db.Create(&user)
	
	userView := view.NewUserView()
	return c.JSON(http.StatusCreated, userView.ToHTTP(user))
}

func (u *UserController) FindUser(c echo.Context) error {
	var user dto.User

	id := c.Param("id")

	if id == "" {
		return c.String(http.StatusBadRequest, "missing id")
	}

	u.Db.Where("id = ?", id).First(&user)

	userView := view.NewUserView()

	return c.JSON(http.StatusOK, userView.ToHTTP(user))
}

func (u *UserController) FetchUsers(c echo.Context) error {
	var users []dto.User

	u.Db.Find(&users)

	var usersDomain []view.UserView
	userView := view.NewUserView()

	for _, v := range users {
		usersDomain = append(usersDomain, userView.ToHTTP(v))
	}
	
	c.JSON(http.StatusOK, usersDomain)

	return nil
} 

func (u *UserController) UpdateUsers(c echo.Context) error {
	var payload dto.UpdateUser

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "missing id")
	}

	if err := c.Bind(&payload); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	payload.ID = id
	var user dto.User
	result := u.Db.First(&user)

	if err := errors.Is(result.Error, gorm.ErrRecordNotFound); err {
		return c.String(http.StatusBadRequest, "user not found")
	}

	if err = c.Validate(payload); err != nil {
		return err
	}

	if payload.Age != 0 {
		user.Age = payload.Age
	}

	if payload.Name != "" {
		user.Name = payload.Name
	}

	if payload.Email != "" {
		user.Email = payload.Email
	}

	u.Db.Save(&user)

	userView := view.NewUserView()
	return c.JSON(http.StatusOK, userView.ToHTTP(user))
}