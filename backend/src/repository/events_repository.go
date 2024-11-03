package repository

import (
	"fmt"
	"myshow/src/config"
	"myshow/src/models"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(cfg *config.Config) (*EventRepository, error) {
	db, err := config.InitDB(cfg)
	if err != nil {
		return nil, err
	}
	return &EventRepository{db: db}, nil
}

func (r *EventRepository) Create(event *models.Event) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var admin models.Admin

		if err := tx.Model(&models.Admin{}).Where("id = ?", event.CreatorID).First(&admin).Error; err != nil {
			return fmt.Errorf("only admins can create events")
		}

		if err := tx.Raw(`
			INSERT INTO events (title, description, location, date, creator_id, artists)
			VALUES (?, ?, ?, ?, ?, ?::text[])
			RETURNING id
		`, event.Title, event.Description, event.Location, event.Date, event.CreatorID, pq.Array(event.Artists)).Scan(&event.ID).Error; err != nil {
			return err
		}

		if err := tx.Model(&admin).Update("events", gorm.Expr("array_append(events, ?)", event.ID)).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *EventRepository) Read() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *EventRepository) ReadByID(id uint) (*models.Event, error) {

	var event models.Event

	result := r.db.First(&event, id)

	if result.Error != nil {

		return nil, result.Error

	}

	return &event, nil

}

func (r *EventRepository) Update(event *models.Event) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var admin models.Admin
		if err := tx.Where("id = ?", event.CreatorID).First(&admin).Error; err != nil {
			return fmt.Errorf("only admins can update events")
		}
		return tx.Save(event).Error
	})
}
func (r *EventRepository) Delete(event *models.Event) error {
	return r.db.Delete(event).Error
}

func (r *EventRepository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
