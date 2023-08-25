package routes

import (
	"github.com/gin-gonic/gin"
)

// AuthRegister route register
func UserRegister(e *gin.Engine) {
	UserRouters := e.Group("/user")
	{
		UserRouters.POST("/login", loginHandler)
		UserRouters.POST("/logout", logoutUserHanler)
		UserRouters.POST("/add")
		// ……
	}
}

// loginHandler
func loginHandler(c *gin.Context) {

}

// logoutUserHanler
func logoutUserHanler(c *gin.Context) {

}
