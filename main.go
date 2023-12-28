package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/account", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, brosephs!",
		})
	})
	router.Run(":8000")
}
