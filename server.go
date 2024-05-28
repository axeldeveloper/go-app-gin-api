package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var users = []user{
	{Id: 546, Username: "John"},
	{Id: 894, Username: "Mary"},
	{Id: 326, Username: "Jane"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8001"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	r.GET("/users", getUsers)

	r.Run(fmt.Sprintf(":%s", port))
}
