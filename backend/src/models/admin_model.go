package models

import (
	"time"

	"github.com/lib/pq"
)

type Admin struct {
	ID        uint          `gorm:"primaryKey"`
	UserID    uint          `gorm:"uniqueIndex;not null" validate:"required"`
	User      uint          `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Events    pq.Int32Array `gorm:"type:int[]" json:"events" validate:"dive,gt=0"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
}
