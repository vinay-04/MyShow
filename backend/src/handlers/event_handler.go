package handlers

import (
	"myshow/src/models"
	"myshow/src/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllEvents(repo *repository.EventRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		events, err := repo.Read()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch events",
			})
		}
		return c.JSON(http.StatusOK, events)
	}
}

func GetEventByID(repo *repository.EventRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid event ID",
			})
		}

		event, err := repo.ReadByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Event not found",
			})
		}
		return c.JSON(http.StatusOK, event)
	}
}

func CreateEvent(repo *repository.EventRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		event := new(models.Event)
		if err := c.Bind(event); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request body",
			})
		}

		if event.Artists == nil {
			event.Artists = make([]string, 0)
		}

		if err := repo.Create(event); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusCreated, event)
	}
}

func UpdateEvent(repo *repository.EventRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid event ID",
			})
		}

		event := new(models.Event)
		if err := c.Bind(event); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request body",
			})
		}

		event.ID = uint(id)
		if err := repo.Update(event); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to update event",
			})
		}
		return c.JSON(http.StatusOK, event)
	}
}

func DeleteEvent(repo *repository.EventRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid event ID",
			})
		}

		event := &models.Event{ID: uint(id)}
		if err := repo.Delete(event); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to delete event",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Event deleted successfully",
		})
	}
}
