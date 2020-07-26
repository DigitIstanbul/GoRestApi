package library

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func SuccessResponse(statusCode int, message string, data interface{}) gin.H {

	if statusCode == 0 {
		statusCode = 200
	}
	return gin.H{"meta": gin.H{"requestId": uuid.New(), "statusCode": http.StatusOK, "message": message}, "result": data}

}
func ErrorResponse(statusCode int, message string) gin.H {

	if statusCode == 0 {
		statusCode = 400
	}
	return gin.H{"meta": gin.H{"requestId": uuid.New(), "statusCode": http.StatusOK, "message": message}, "result": gin.H{}}

}
