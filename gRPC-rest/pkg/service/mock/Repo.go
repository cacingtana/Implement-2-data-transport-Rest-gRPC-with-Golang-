package mocks

import (
	"grpc-rest/pkg/models"

	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (t *Repository) InsertLog(_t models.Logs) error {
	ret := t.Called(_t)

	var tProductError error
	if rf, ok := ret.Get(0).(func(models.Logs) error); ok {
		tProductError = rf(_t)
	} else {
		tProductError = ret.Error(0)
	}
	return tProductError
}
