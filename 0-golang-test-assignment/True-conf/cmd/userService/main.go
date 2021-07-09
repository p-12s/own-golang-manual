package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log2 "github.com/labstack/gommon/log"
	"os"
	"userService/internal/pkg/user/delivery"
	"userService/internal/pkg/user/repository"
	"userService/internal/pkg/user/usecase"
	"userService/internal/userService"
)

func setConfig(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Logger.SetLevel(log2.Lvl(0))

	tableName := os.Getenv("TableName")
	dir := os.Getenv("DataDir")

	db := userService.InitDB(tableName, dir)
	userRepository := repository.NewUserRepository(*db, tableName)
	userUsecase := usecase.UserUsecase{Repository: userRepository}
	userDelivery := delivery.UserDelivery{Usecase: &userUsecase}
	userDelivery.SetRoutersForUser(e)
}

func main() {
	e := echo.New()
	setConfig(e)
	e.Logger.Fatal(e.Start(":8080"))
}
