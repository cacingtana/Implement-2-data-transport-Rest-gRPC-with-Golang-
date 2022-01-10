package service

import (
	"grpc-rest/pkg/models"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) InsertLog(logs models.Logs) error {

	err := s.repository.InsertLog(logs)
	if err != nil {
		return err
	}
	return nil

}
