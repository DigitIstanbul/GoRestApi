package main

import (
	"BudgetApi/src/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1.0")
	{
		v1.GET("/", handler.Index)
		v1.POST("/user/register", handler.Register)
		v1.POST("/user/login", handler.Login)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "NO_METHOD", "message": "Page not found"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
