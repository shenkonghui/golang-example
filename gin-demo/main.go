package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": os.Getenv("version"),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
