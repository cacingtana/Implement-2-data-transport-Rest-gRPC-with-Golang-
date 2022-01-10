package response

import "net/http"

type SuccessResponseCode string

const (
	Success SuccessResponseCode = "success"
)

type SuccessResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
}

func NewSuccessResponse(data interface{}, msg string) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		msg,
		data,
	}
}

func NewSuccessResponseWithoutData(msg string) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		msg,
		map[string]interface{}{},
	}
}

//NewSuccessResponse create new success payload
func NewSuccessResponseNoContent(msg string) (int, SuccessResponse) {
	return http.StatusNoContent, SuccessResponse{
		Success,
		msg,
		map[string]interface{}{},
	}
}
