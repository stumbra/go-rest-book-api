package main

import (
	"go-rest-book-api/controllers"
	_ "go-rest-book-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:9090
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	v1 := router.Group("/api/v1")

	{
		books := v1.Group("/books")
		{
			books.GET("", controllers.GetBooks)
			books.POST("", controllers.AddBook)
			books.GET(":id", controllers.GetBook)
			books.PATCH(":id", controllers.UpdateBookAvailability)
			books.DELETE(":id", controllers.RemoveBook)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:9090")
}
