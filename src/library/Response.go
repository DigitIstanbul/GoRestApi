package library

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ResponseStruct struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	RequestId  uuid.UUID `json:"requestId"`
	Response   gin.H     `json:"requestId"`
}

func SuccessResponse(statusCode int, message string, data interface{}) ResponseStruct {

	requestId := uuid.New()
	response := gin.H{"meta": gin.H{"requestId": requestId, "statusCode": http.StatusOK, "message": message}, "result": data}
	return ResponseStruct{Message: message, StatusCode: statusCode, Response: response, RequestId: requestId}

}

func ErrorResponse(statusCode int, message string, data interface{}) ResponseStruct {

	if statusCode == 0 {
		statusCode = 400
	}

	requestId := uuid.New()
	response := gin.H{"meta": gin.H{"requestId": requestId, "statusCode": http.StatusOK, "message": message}, "result": data}
	return ResponseStruct{Message: message, StatusCode: statusCode, Response: response, RequestId: requestId}

}
