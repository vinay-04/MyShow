package config

import (
	"fmt"
	"myshow/src/models"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: import from .Env

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Port       string
	JWTSecret  string
}

func getEnvOrDefault(envKey, defaultValue string) string {
	if value, exists := os.LookupEnv(envKey); exists {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	return &Config{
		DBHost:     getEnvOrDefault("DB_HOST", "localhost"),
		DBPort:     getEnvOrDefault("DB_PORT", "5432"),
		DBUser:     getEnvOrDefault("DB_USER", "postgres"),
		DBPassword: getEnvOrDefault("DB_PASSWORD", "admin1234"),
		DBName:     getEnvOrDefault("DB_NAME", "myshow"),
		Port:       getEnvOrDefault("PORT", "8080"),
		JWTSecret:  getEnvOrDefault("JWT_SECRET", "myshow-secret"),
	}, nil
}

func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	pgConfig := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}

	db, err := gorm.Open(postgres.New(pgConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Event{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate: %v", err)
	}

	return db, nil
}
