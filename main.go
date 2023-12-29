package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string
	Name string
}

var users []*User

func main() {
	router := gin.Default()
	users = append(users, &User{Id: "1", Name: "vijendra"})
	router.GET("/", getAllUsers)
	router.Run(":8000")
}

func getAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}
