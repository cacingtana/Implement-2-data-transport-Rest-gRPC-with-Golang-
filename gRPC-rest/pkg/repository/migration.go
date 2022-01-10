package repository

import (
	"grpc-rest/pkg/models"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Logs{})
}
