package middleware

import (
	"BudgetApi/src/library"
	"github.com/gin-gonic/gin"
	"os"
)

func Auth(c *gin.Context) {
	if c.GetHeader("Auth") == os.Getenv("REST_AUTH_KEY") {
		c.Next()
	} else {
		c.JSON(401, library.ErrorResponse(401, "Auth error!"))
		c.Abort()
	}
}
