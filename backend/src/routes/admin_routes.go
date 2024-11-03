package routes

import (
	"myshow/src/handlers"
	"myshow/src/middleware"
	"myshow/src/repository"

	"github.com/labstack/echo/v4"
)

func SetupAdminRoutes(e *echo.Echo, adminRepo *repository.AdminRepository, jwtSecret string) {
	admin := e.Group("/api/admin")
	admin.Use(middleware.JWTMiddleware(jwtSecret))
	admin.Use(middleware.AdminOnly)

	admin.GET("", handlers.GetAllAdmins(*adminRepo))
	admin.GET("/:id", handlers.GetAdminByUsername(*adminRepo))
}
