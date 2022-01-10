package mocks

import (
	"grpc-rest/pkg/models"

	mock "github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (t *Service) InsertLog(logs models.Logs) error {
	ret := t.Called(logs)

	var tProductError error
	if rf, ok := ret.Get(0).(func(models.Logs) error); ok {
		tProductError = rf(logs)
	} else {
		tProductError = ret.Error(0)
	}

	return tProductError
}
