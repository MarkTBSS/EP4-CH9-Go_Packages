package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersHandler(echoContext echo.Context) error {
	stagement, err := database.Prepare("SELECT id, name, age FROM users")
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, Err{Message: "Can't prepare query all users statment:" + err.Error()})
	}
	rowResults, err := stagement.Query()
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, Err{Message: "Can't query all users:" + err.Error()})
	}
	userStructArray := []User{}
	for rowResults.Next() {
		userStruct := User{}
		err := rowResults.Scan(&userStruct.ID, &userStruct.Name, &userStruct.Age)
		if err != nil {
			return echoContext.JSON(http.StatusInternalServerError, Err{Message: "Can't scan user:" + err.Error()})
		}
		userStructArray = append(userStructArray, userStruct)
	}
	return echoContext.JSON(http.StatusOK, userStructArray)
}

// Step 5 : getUserHandler function
func GetUserHandler(echoContext echo.Context) error {
	id := echoContext.Param("id")
	stagement, err := database.Prepare("SELECT id, name, age FROM users WHERE id = $1")
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query user statment:" + err.Error()})
	}
	rowResult := stagement.QueryRow(id)
	userStruct := User{}
	err = rowResult.Scan(&userStruct.ID, &userStruct.Name, &userStruct.Age)
	switch err {
	case sql.ErrNoRows:
		return echoContext.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return echoContext.JSON(http.StatusOK, userStruct)
	default:
		return echoContext.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}
}
