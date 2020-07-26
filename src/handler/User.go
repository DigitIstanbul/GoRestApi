package handler

import (
	"BudgetApi/src/library"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	response := library.SuccessResponse(http.StatusAccepted, "Success", nil)
	c.JSON(response.StatusCode, response.Response)

}

func Register(c *gin.Context) {

	response := library.SuccessResponse(http.StatusCreated, "Success", nil)
	c.JSON(response.StatusCode, response.Response)

}
