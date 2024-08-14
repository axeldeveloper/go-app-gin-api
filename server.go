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
	// "github.com/swaggo/files"
	// "github.com/swaggo/gin-swagger"

	docs "example-go-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "Hello, world!", 
		"status": true,
	})
}




// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8001"
	//docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	
	port := os.Getenv("PORT")

	if port == "" {
		port = "8001"
	}

	r := gin.Default()

	database := db.Connection()
	// Handlers
	companyRepository := repository.NewUserRepository(database)
	companyHandler := handlers.NewUserHandler(companyRepository)

	
	r.GET("/", Helloworld)
	
	r.GET("/ping", func(c *gin.Context) {
		err := godotenv.Load() //by default, it is .env so we don't have to write
		if err != nil {
			c.JSON(404, gin.H{
				"Error": "Error is occurred  on .env file please check",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message":     "pong",
			"DB_HOST":     os.Getenv("DB_HOST"),
			"ENV_STATION": os.Getenv("ENV_STATION"),
		})

	})

	v1 := r.Group("/api/v1")
	{
		u := v1.Group("/users")
		{
			u.GET("/", companyHandler.GetCompanies )

			u.POST("/", companyHandler.CreateUser)
			u.GET("/users/:id", companyHandler.GetUserByID)
			u.PUT("/users/:id", companyHandler.UpdateUser)
			u.DELETE("/users/:id", companyHandler.DeleteUser)
		}


		c := v1.Group("/companies")
		{
			c.GET("/", companyHandler.GetCompanies)
		}
	}
   	
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(fmt.Sprintf(":%s", port))
}
