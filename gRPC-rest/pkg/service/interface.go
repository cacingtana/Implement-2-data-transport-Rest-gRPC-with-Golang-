package service

import (
	"grpc-rest/pkg/models"
)

type Service interface {
	InsertLog(moviesLog models.Logs) error
}

type Repository interface {
	InsertLog(moviesLog models.Logs) error
}
