package models

import (
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"uniqueIndex;not null" validate:"required,min=3,max=32"`
	Password  string         `gorm:"not null" validate:"required,min=8"`
	FirstName string         `gorm:"not nul" validate:"required,alpha"`
	LastName  string         `gorm:"not null" validate:"required,alpha"`
	Email     string         `gorm:"uniqueIndex;not null" validate:"required,email"`
	Phone     string         `gorm:"uniqueIndex;not null" validate:"required,e164"`
	Admin     bool           `gorm:"not null;default:false"`
	Events    pq.StringArray `gorm:"type:text[]" json:"events"`
	AdminData *Admin         `gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
}

func (u *User) IsAdmin() bool {
	return u.Admin
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	if u.Admin {
		admin := Admin{
			UserID: u.ID,
		}
		if err := tx.Create(&admin).Error; err != nil {
			return err
		}
	}
	return nil
}
