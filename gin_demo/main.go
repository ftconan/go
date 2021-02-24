// Author: magician
// File:   main.go
// Date:   2021/2/24
package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default() // Logger, Recovery (middleware)
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "ft")
		pwd := c.Query("pwd")

		c.JSON(200, gin.H{
			"id": id,
			"user": user,
			"pwd": pwd,
			"method": c.Request.Method,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "ft")
		pwd := c.PostForm("pwd")

		c.JSON(200, gin.H{
			"id": 1,
			"user": user,
			"pwd": pwd,
			"method": c.Request.Method,
		})
	})
	r.PUT("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultPostForm("user", "ft")
		pwd := c.PostForm("pwd")

		c.JSON(200, gin.H{
			"id": id,
			"user": user,
			"pwd": pwd,
			"method": c.Request.Method,
		})
	})
	r.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(200, gin.H{
			"id": id,
			"method": c.Request.Method,
		})
	})

	//r.Run(":1010")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
