package repository

import (
	"grpc-rest/pkg/models"

	"gorm.io/gorm"
)

type DbRepo struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *DbRepo {
	return &DbRepo{
		db,
	}
}

func (d *DbRepo) InsertLog(logs models.Logs) error {
	err := d.db.Save(logs).Error
	if err != nil {
		return err
	}
	return nil
}
