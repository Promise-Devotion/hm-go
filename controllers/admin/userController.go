package admin

import (
	"mygin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (con AdminController) UserIndex(c *gin.Context) {

	userList := []models.User{}

	models.DB.Find(&userList)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    userList,
		"message": "查询成功",
	})
}
