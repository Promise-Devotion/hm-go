package admin

import (
	"fmt"
	"mygin/models"
	"net/http"
	"strconv"

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

func (con AdminController) UserOne(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Request.FormValue("id"))

	if err == nil {
		fmt.Println(user_id)

		userResult := models.User{UserId: user_id}

		models.DB.Where("user_id=?", user_id).Find(&userResult)

		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"data":    userResult,
			"message": "查询成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"data":    err,
			"message": "查询失败",
		})
	}

}
