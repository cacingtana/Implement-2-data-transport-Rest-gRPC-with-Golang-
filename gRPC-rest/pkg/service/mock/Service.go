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

	var tLogs error
	if rf, ok := ret.Get(0).(func(models.Logs) error); ok {
		tLogs = rf(logs)
	} else {
		tLogs = ret.Error(0)
	}

	return tLogs
}
