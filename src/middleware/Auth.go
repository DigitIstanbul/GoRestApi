package middleware

import (
	"BudgetApi/src/config"
	"BudgetApi/src/library"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	if c.GetHeader("Auth") == config.GetConfig("REST_AUTH_KEY") {
		c.Next()
	} else {

		response := library.ErrorResponse(http.StatusForbidden, "Auth error!", nil)
		c.JSON(response.StatusCode, response.Response)
		c.Abort()
	}
}
