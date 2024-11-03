package handlers

import (
	"myshow/src/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllAdmins(repo repository.AdminRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		admins, err := repo.Read()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch admins",
			})
		}
		return c.JSON(http.StatusOK, admins)
	}
}

func GetAdminByUsername(repo repository.AdminRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		admin, err := repo.ReadByUsername(username)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Admin not found",
			})
		}
		return c.JSON(http.StatusOK, admin)
	}
}
