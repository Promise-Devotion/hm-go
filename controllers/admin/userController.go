package admin

import "github.com/gin-gonic/gin"

type AdminController struct{}

func (con AdminController) UserIndex(c *gin.Context) {
	c.JSON(200, "user list")
}
