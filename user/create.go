package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserHandler(echoContext echo.Context) error {
	// Step 3.1
	userStruct := User{}
	err := echoContext.Bind(&userStruct)
	if err != nil {
		// Step 3.2
		return echoContext.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	row := database.QueryRow("INSERT INTO users (name, age) values ($1, $2)  RETURNING id", userStruct.Name, userStruct.Age)
	err = row.Scan(&userStruct.ID)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return echoContext.JSON(http.StatusCreated, userStruct)
}
