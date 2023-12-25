package server

import (
	"github.com/labstack/echo/v4"
)

func CheckUserRoleHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("User-Role") == "admin" {
			c.Logger().Info("admin user")
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
