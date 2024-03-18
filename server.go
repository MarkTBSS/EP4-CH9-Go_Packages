package main

import (
	"log"

	"github.com/MarkTBSS/EP4-CH9-Go_Packages/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	user.InitDB()
	echoInstance := echo.New()
	echoInstance.Use(middleware.BasicAuth(user.Login))
	//echoInstance.Use(middleware.Logger())
	//echoInstance.Use(middleware.Recover())
	echoInstance.POST("/users", user.CreateUserHandler)
	echoInstance.GET("/users", user.GetUsersHandler)
	echoInstance.GET("/users/:id", user.GetUserHandler)
	log.Fatal(echoInstance.Start(":2567"))
}
