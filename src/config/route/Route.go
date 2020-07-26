package route

import (
	"BudgetApi/src/config"
	"BudgetApi/src/handler"
	"BudgetApi/src/library"
	"BudgetApi/src/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Routes struct {
}

func (c Routes) StartGin() {
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

	_ = r.Run(":" + config.GetConfig("PORT"))
}
