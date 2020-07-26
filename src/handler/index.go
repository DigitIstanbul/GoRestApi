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
	c.JSON(http.StatusOK, library.SuccessResponse(http.StatusOK, "Success", user))
}
