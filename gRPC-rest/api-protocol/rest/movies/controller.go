package movies

import (
	"fmt"
	"grpc-rest/pkg/models"
	"grpc-rest/pkg/service"
	"time"

	"grpc-rest/api-protocol/rest/request"
	"grpc-rest/api-protocol/rest/response"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service service.Service
	request request.RequestMovies
}

func NewController(service service.Service) *Controller {
	return &Controller{
		service,
		request.RequestMovies{},
	}
}

func (cont *Controller) GetMovies(c echo.Context) error {

	pagination := c.QueryParam("page")
	searchWord := c.QueryParam("s")
	result := cont.request.GetAllMovies(pagination, searchWord)

	method := "GET"
	params := fmt.Sprintf("Params : %s/%s", searchWord, pagination)
	modelLog := models.NewLogs(
		method,
		params,
		time.Now().UTC(),
	)
	cont.service.InsertLog(modelLog)
	return c.JSON(response.NewSuccessResponse(result, "successfully fetch all movies"))
}

func (cont *Controller) GetMoviesById(c echo.Context) error {
	id := c.QueryParam("i")
	result := cont.request.GetMoviesById(id)

	method := "GET"
	params := fmt.Sprintf("Params : %s", id)
	modelLog := models.NewLogs(
		method,
		params,
		time.Now().UTC(),
	)
	cont.service.InsertLog(modelLog)
	return c.JSON(response.NewSuccessResponse(result, "successfully fetch movie by ID"))
}
