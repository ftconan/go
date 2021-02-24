// Author: magician
// File:   main.go
// Date:   2021/2/24
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	Name string `json:"name" uri:"name" form:"name"`
	Age  int    `json:"age" uri:"age" form:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex" uri:"sex" form:"sex"`
}

func mustBig(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(int) <= 18 {
		return false
	}
	return true
}

func main() {
	r := gin.Default() // Logger, Recovery (middleware)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mustBig", mustBig)
	}

	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "ft")
		pwd := c.Query("pwd")

		c.JSON(200, gin.H{
			"id":     id,
			"user":   user,
			"pwd":    pwd,
			"method": c.Request.Method,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "ft")
		pwd := c.PostForm("pwd")

		c.JSON(200, gin.H{
			"id":     1,
			"user":   user,
			"pwd":    pwd,
			"method": c.Request.Method,
		})
	})
	r.PUT("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultPostForm("user", "ft")
		pwd := c.PostForm("pwd")

		c.JSON(200, gin.H{
			"id":     id,
			"user":   user,
			"pwd":    pwd,
			"method": c.Request.Method,
		})
	})
	r.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(200, gin.H{
			"id":     id,
			"method": c.Request.Method,
		})
	})
	r.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		// ShouldBind 通过content-type判断
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(400, gin.H{
				"msg":  "报错了",
				"data": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功了",
				"data": p,
			})
		}
	})

	//r.Run(":8080")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
