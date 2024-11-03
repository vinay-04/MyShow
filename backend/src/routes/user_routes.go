package routes

import (
	"myshow/src/handlers"
	"myshow/src/middleware"
	"myshow/src/repository"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo, userRepo *repository.UserRepository, eventRepo *repository.EventRepository, jwtSecret string) {
	users := e.Group("/api/users")

	users.POST("/register", handlers.RegisterUser(userRepo))
	users.POST("/login", handlers.LoginUser(userRepo))

	protected := users.Group("")
	protected.GET("", handlers.GetAllUsers(userRepo))
	protected.Use(middleware.JWTMiddleware(jwtSecret))
	protected.GET("/:username", handlers.GetUserByUsername(userRepo))
	protected.PUT("/:username", handlers.UpdateUser(userRepo))
	protected.DELETE("/:username", handlers.DeleteUser(userRepo))
	protected.PUT("/mark-event", handlers.AddUserToEvent(userRepo, eventRepo))
}
