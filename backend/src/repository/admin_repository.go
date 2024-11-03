package repository

import (
	"myshow/src/config"
	"myshow/src/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(cfg *config.Config) (*AdminRepository, error) {
	db, err := config.InitDB(cfg)
	if err != nil {
		return nil, err
	}

	return &AdminRepository{
		db: db,
	}, nil
}

func (r *AdminRepository) Create(admin *models.Admin) error {
	return r.db.Create(admin).Error
}

func (r *AdminRepository) Read() ([]models.Admin, error) {
	var admins []models.Admin
	err := r.db.Find(&admins).Error
	return admins, err
}

func (r *AdminRepository) ReadByUsername(username string) (models.Admin, error) {
	var admin models.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return admin, err
}

func (r *AdminRepository) Update(admin *models.Admin) error {
	return r.db.Save(admin).Error
}

func (r *AdminRepository) Delete(username string) error {
	return r.db.Where("username = ?", username).Delete(&models.Admin{}).Error
}

func (r *AdminRepository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
