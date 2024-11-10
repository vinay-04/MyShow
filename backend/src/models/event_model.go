package models

import (
	"time"

	"github.com/lib/pq"
)

type Event struct {
	ID          uint           `gorm:"primaryKey"`
	Title       string         `gorm:"not null" validate:"required,min=3,max=100"`
	ImageURL    string         `gorm:"not null" validate:"required,url"`
	Description string         `gorm:"not null" validate:"required,min=10,max=500"`
	Location    string         `gorm:"not null" validate:"required"`
	Date        time.Time      `gorm:"not null" validate:"required,gtefield=CreatedAt"`
	CreatorID   uint           `gorm:"foreignKey:creator_id" validate:"required"`
	Artists     pq.StringArray `gorm:"type:text[]" json:"artists" validate:"unique"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
}
