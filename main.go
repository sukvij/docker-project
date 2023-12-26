package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/users", getAllUsers)
	app.Run(":8000")
}

func getAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "welcome")
}
