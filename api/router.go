package api

import (
	v1 "FINE/api/handler/example/v1"
	"FINE/api/middleware"
	"github.com/gin-gonic/gin"
)

func RouterHandler(router *gin.Engine) {
	set(router)
}

func set(router *gin.Engine) {
	if router == nil {
		return
	}

	example := router.Group("/example").Use(middleware.Base(middleware.InputIn{}), middleware.ResponseJSON)
	{
		example.GET("/hello", v1.Hello)
	}
}
