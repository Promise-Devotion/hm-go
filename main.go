package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"data":    "pong",
		})
	})
	r.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}