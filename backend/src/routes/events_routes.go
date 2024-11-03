package routes

import (
	"myshow/src/handlers"
	"myshow/src/middleware"
	"myshow/src/repository"

	"github.com/labstack/echo/v4"
)

func SetupEventRoutes(e *echo.Echo, eventRepo *repository.EventRepository, jwtSecret string) {
	events := e.Group("/api/events")

	events.GET("", handlers.GetAllEvents(eventRepo))
	events.GET("/:id", handlers.GetEventByID(eventRepo))

	adminEvents := events.Group("")
	adminEvents.Use(middleware.AdminOnly)
	adminEvents.POST("/create", handlers.CreateEvent(eventRepo))
	adminEvents.PUT("/:id", handlers.UpdateEvent(eventRepo))
	adminEvents.DELETE("/:id", handlers.DeleteEvent(eventRepo))
}
