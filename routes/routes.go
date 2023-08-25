package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter 配置路由
func SetupRouter() *gin.Engine {
	route := gin.Default()
	route.Use(cors.Default())
	// other config

	// register all route.
	UserRegister(route)
	BillRegister(route)
	PingRegister(route)
	// ……

	return route
}
