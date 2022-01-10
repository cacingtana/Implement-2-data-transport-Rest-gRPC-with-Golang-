package service_test

import (
	"grpc-rest/pkg/models"
	"grpc-rest/pkg/service"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	logsMock "grpc-rest/pkg/service/mock"
)

var (
	id       = 1
	Method   = "GET"
	Params   = "abcde"
	accessAt = "time.Time"
)

var (
	logService service.Service
	logRepo    logsMock.Repository

	logssData []models.Logs = make([]models.Logs, 0)

	logData   models.Logs
	createLog models.Logs
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {

	logData = models.NewLogs(
		Method,
		Params,
		time.Now().UTC(),
	)

	logService = service.NewService(&logRepo)

}

func TestInsertLogs(t *testing.T) {
	t.Run("Expect insert log success", func(t *testing.T) {
		logRepo.On("InsertLog", mock.AnythingOfType("models.Logs"), mock.AnythingOfType("string")).Return(nil).Once()

		err := logService.InsertLog(logData)
		assert.Nil(t, err)

	})

	t.Run("Expect insert log failed", func(t *testing.T) {
		logRepo.On("CreateProduct", mock.AnythingOfType("models.Logs"), mock.AnythingOfType("string")).Return("ErrInternalServerError").Once()

		err := logService.InsertLog(logData)
		assert.NotNil(t, err)
		assert.Equal(t, err, "ErrInternalServerError")
	})
}

//Belum selesai
