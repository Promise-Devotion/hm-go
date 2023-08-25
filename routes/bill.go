package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthRegister route register
func BillRegister(e *gin.Engine) {
	BillRouters := e.Group("/bill")
	{
		BillRouters.GET("/bill-list", billListHandler)
		BillRouters.POST("/bill-task", billTaskHandler)
		// ……
	}
}

// loginHandler
func billListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    "this is bill-list",
		"message": "操作成功",
	})
}

// logoutUserHanler
func billTaskHandler(c *gin.Context) {

}
