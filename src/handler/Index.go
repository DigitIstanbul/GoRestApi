package handler

import (
	"BudgetApi/src/library"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func Index(c *gin.Context) {

	user := &User{Name: "Muhammed Ali"}

	response := library.SuccessResponse(http.StatusOK, "Success", user)
	c.JSON(response.StatusCode, response.Response)

}
