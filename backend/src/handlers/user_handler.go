package handlers

import (
	"myshow/src/middleware"
	"myshow/src/models"
	"myshow/src/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterUser(repo *repository.UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request body",
			})
		}

		if err := user.HashPassword(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Error hashing password",
			})
		}

		if err := repo.Create(user); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		token, err := middleware.GenerateToken(user.ID, user.Username, user.Admin, c.Get("jwt_secret").(string))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to generate token",
			})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"token": token,
			"user":  user})
	}
}

func LoginUser(repo *repository.UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		credentials := struct {
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		}{}

		if err := c.Bind(&credentials); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid credentials",
			})
		}

		user, err := repo.ReadByUsername(credentials.Username)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid credentials",
			})
		}

		if !user.ComparePassword(credentials.Password) {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid credentials",
			})
		}

		token, err := middleware.GenerateToken(user.ID, user.Username, user.Admin, c.Get("jwt_secret").(string))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to generate token",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
			"user":  user,
		})
	}
}

func AddUserToEvent(userRepo *repository.UserRepository, eventRepo *repository.EventRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		payload := struct {
			Username string `json:"username" validate:"required"`
			EventID  uint   `json:"event_id" validate:"required"`
		}{}

		if err := c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request body",
			})
		}

		event, err := eventRepo.ReadByID(payload.EventID)
		user, _ := userRepo.ReadByUsername(payload.Username)

		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Event not found",
			})
		}
		user.Events = append(user.Events, strconv.FormatUint(uint64(payload.EventID), 10))

		event.Artists = append(event.Artists, payload.Username)

		if err := userRepo.Update(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update user",
			})
		}

		if err := eventRepo.Update(event); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update event",
			})
		}

		return c.JSON(http.StatusOK, event)
	}
}

func GetAllUsers(repo *repository.UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := repo.Read()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch users",
			})
		}
		return c.JSON(http.StatusOK, users)
	}
}

func GetUserByUsername(repo *repository.UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		user, err := repo.ReadByUsername(username)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "User not found",
			})
		}
		return c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(repo *repository.UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		existingUser, err := repo.ReadByUsername(username)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "User not found",
			})
		}

		updatedUser := new(models.User)
		if err := c.Bind(updatedUser); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request body",
			})
		}

		updatedUser.ID = existingUser.ID
		if err := repo.Update(updatedUser); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update user",
			})
		}

		return c.JSON(http.StatusOK, updatedUser)
	}
}

func DeleteUser(repo *repository.UserRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		if err := repo.Delete(username); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to delete user",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "User deleted successfully",
		})
	}
}
