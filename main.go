package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string
	Name string
}

var users map[string]*User

func main() {
	router := gin.Default()
	users = make(map[string]*User)
	users["1"] = &User{Id: "1", Name: "vijendra"}
	users["2"] = &User{Id: "2", Name: "sukariya"}
	router.GET("/", getAllUsers)
	router.GET("/{id:id}", getUserById)
	router.Run(":8000")
}

func getAllUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, users)
}

func getUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	ctx.JSON(http.StatusOK, users[id])
}
