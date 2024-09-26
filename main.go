package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/lucasmbrute2/go-api/internal/http/controllers"
	"github.com/lucasmbrute2/go-api/internal/modules/user/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func main(){
	e := echo.New()

	dsn := "root:root@tcp(127.0.0.1)/dev1"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.UserEntity{})
	if err != nil{
		panic(err)
	}
	e.Validator = &CustomValidator{ validator: validator.New() }

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	usersGroup := e.Group("/users")

	userController := controllers.NewUserController(db)
	usersGroup.POST("/", 	userController.CreateUser)
	usersGroup.GET("/:id", 	userController.FindUser)

	e.Logger.Fatal(e.Start(":3001"))
}