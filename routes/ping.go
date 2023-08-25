package routes

import (
	"mygin/controllers/admin"

	"github.com/gin-gonic/gin"
)

// AuthRegister route register
func PingRegister(e *gin.Engine) {
	PingRouters := e.Group("/ping")
	{
		PingRouters.GET("/demo", admin.PingController{}.PingDemo)
		// ……
	}
}
