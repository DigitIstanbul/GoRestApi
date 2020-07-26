package main

import (
	"BudgetApi/src/config"
	"BudgetApi/src/handler"
	"BudgetApi/src/library"
	"BudgetApi/src/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	config.LoadEnv()
	r := gin.Default()

	v1 := r.Group("/v1.0")
	v1.Use(middleware.Auth)
	{
		v1.GET("/", handler.Index)
		userGroup := v1.Group("user")
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
	}

	r.NoRoute(func(c *gin.Context) {
		response := library.ErrorResponse(http.StatusNotFound, "Service not found!", nil)
		c.JSON(response.StatusCode, response.Response)
	})

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
