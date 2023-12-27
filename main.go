package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	router := gin.Default()

	router.GET("/user", getAllUser)
	router.Run(":8000")
}

func getAllUser(ctx *gin.Context) {
	ctx.JSON(200, "great")
}
