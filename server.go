package main

import (
	"example-go-gin/db"
	"example-go-gin/handlers"
	"example-go-gin/repository"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	r.GET("/ping", func(c *gin.Context) {
		err := godotenv.Load() //by default, it is .env so we don't have to write
		if err != nil {
			c.JSON(404, gin.H{
				"Error": "Error is occurred  on .env file please check",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"DB_HOST": os.Getenv("DB_HOST"),
		})

	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", name),
		})
	})

	database := db.Connection()
	// Handlers
	companyRepository := repository.NewUserRepository(database)
	companyHandler := handlers.NewUserHandler(companyRepository)

	r.GET("/companies/", companyHandler.GetCompanies)
	r.POST("/users", companyHandler.CreateUser)
	r.GET("/users/:id", companyHandler.GetUserByID)
	r.PUT("/users/:id", companyHandler.UpdateUser)
	r.DELETE("/users/:id", companyHandler.DeleteUser)

	r.Run(fmt.Sprintf(":%s", port))
}
