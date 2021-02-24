// Author: magician
// File:   main.go
// Date:   2021/2/24
package main
import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default() // Logger, Recovery (middleware)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
