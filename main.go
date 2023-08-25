package main

import (
	"mygin/routes"
)

func main() {
	// r := gin.Default()
	r := routes.SetupRouter()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code":    200,
	// 		"message": "success",
	// 		"data":    "pong",
	// 	})
	// })
	r.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
