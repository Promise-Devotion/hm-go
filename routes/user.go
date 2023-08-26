package routes

import (
	"mygin/controllers/admin"

	"github.com/gin-gonic/gin"
)

// AuthRegister route register
func UserRegister(e *gin.Engine) {
	UserRouters := e.Group("/user")
	{
		UserRouters.POST("/login", loginHandler)
		UserRouters.POST("/logout", logoutUserHanler)
		UserRouters.GET("/userlist", admin.AdminController{}.UserIndex)
		UserRouters.GET("/getuser", admin.AdminController{}.UserOne)
		// ……
	}
}

// loginHandler
func loginHandler(c *gin.Context) {

}

// logoutUserHanler
func logoutUserHanler(c *gin.Context) {

}
