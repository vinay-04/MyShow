package main

import (
	"log"
	"myshow/src/config"
	"myshow/src/middleware"
	"myshow/src/repository"
	"myshow/src/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("jwt_secret", cfg.JWTSecret)
			return next(c)
		}
	})

	userRepo, err := repository.NewUserRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	adminRepo, err := repository.NewAdminRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}

	eventRepo, err := repository.NewEventRepository(cfg)
	if err != nil {
		log.Fatal(err)
	}
	e.Use(middleware.Validate)

	e.Use(echo.WrapMiddleware(middleware.LoggingMiddleware))

	routes.SetupUserRoutes(e, userRepo, eventRepo, cfg.JWTSecret)
	routes.SetupAdminRoutes(e, adminRepo, cfg.JWTSecret)
	routes.SetupEventRoutes(e, eventRepo, cfg.JWTSecret)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to myshow API")
	})

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
