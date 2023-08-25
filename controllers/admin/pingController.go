package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (con PingController) PingDemo(c *gin.Context) {
	userid := c.Request.FormValue("id")
	fmt.Println(userid + "------------------")
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    "this is bill-list",
		"message": "操作成功",
	})
}
