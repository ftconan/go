// Author: magician
// File:   main.go
// Date:   2021/2/24
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
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

func middle1()gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println("middle1 before")
		c.Next()
		fmt.Println("middle1 last")
	}
}

func middle2()gin.HandlerFunc{
	return func(c *gin.Context) {
		fmt.Println("middle2 before")
		c.Next()
		fmt.Println("middle2 last")
	}
}

func main() {
	r := gin.Default() // Logger, Recovery (middleware)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mustBig", mustBig)
	}

	// Using GET, POST, PUT, PATCH, DELETE and OPTIONS
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

	// Model binding and validation
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

	// Upload files
	r.POST("/testUpload", func(c *gin.Context) {
		// Single file
		//file, _ := c.FormFile("file")
		//name := c.PostForm("name")
		//dst := "./"+file.Filename
		// Upload the file to specific dst.
		//_ = c.SaveUploadedFile(file, dst)

		// os
		//src, _ := file.Open()
		//defer src.Close()
		//out, _ := os.Create(dst)
		//defer out.Close()
		//_, _ = io.Copy(out, src)

		//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
		//fmt.Println(file.Filename)
		//c.File(dst)

		//c.JSON(200, gin.H{
		//	"msg": file,
		//	"name": name,
		//})

		// Multiple file
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for _, file:= range files {
			log.Println(file.Filename)
			dst := "./"+file.Filename
			// Upload the file to specific dst.
			_ = c.SaveUploadedFile(file, dst)
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
			c.File(dst)
		}
	})

	// Grouping routes
	v1 := r.Group("/v1")
	// Using middleware
	v1.Use(middle1()).Use(middle2())

	v1.GET("test", func(c *gin.Context) {
		fmt.Println("test method")
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
