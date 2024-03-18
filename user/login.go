package user

import "github.com/labstack/echo/v4"

func Login(username, password string, c echo.Context) (bool, error) {
	if username == "mark" && password == "12345" {
		return true, nil
	}
	return false, nil
}
