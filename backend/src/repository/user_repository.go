package repository

import (
	"myshow/src/config"
	"myshow/src/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(cfg *config.Config) (*UserRepository, error) {
	db, err := config.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	return &UserRepository{
		db: db,
	}, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *UserRepository) Read() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) ReadByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) ReadByID(id uint) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var oldUser models.User
		if err := tx.First(&oldUser, user.ID).Error; err != nil {
			return err
		}

		if user.Password != oldUser.Password {
			if err := user.HashPassword(); err != nil {
				return err
			}
		}
		if oldUser.Admin != user.Admin {
			if user.Admin {
				admin := &models.Admin{UserID: user.ID}
				if err := tx.Create(admin).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Where("user_id = ?", user.ID).Delete(&models.Admin{}).Error; err != nil {
					return err
				}
			}
		}

		return tx.Save(user).Error
	})
}

func (r *UserRepository) Delete(username string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var user models.User
		if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
			return err
		}
		if err := tx.Table("event_artists").Where("user_id = ?", user.ID).Delete(nil).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", user.ID).Delete(&models.Admin{}).Error; err != nil {
			return err
		}
		return tx.Delete(&user).Error
	})
}

func (r *UserRepository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
