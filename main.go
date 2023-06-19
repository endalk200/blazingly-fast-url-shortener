package main

import (
	"fmt"

	"github.com/endalk200/blazingly-fast-url-shortener/docs"
	handler "github.com/endalk200/blazingly-fast-url-shortener/handler"
	"github.com/endalk200/blazingly-fast-url-shortener/store"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Blazingly Fast URL Shortener API
// @version         1.0
// @description     Blazingly Fast URL Shortener API written in Go using Gin framework.

// @contact.name   endalk200
// @contact.email  eb808826@gmail.com

// @host      localhost:9808
// @BasePath  /
func main() {
	gin.ForceConsoleColor()

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	// Root route for getting API information
	// @summary Get API information
	// @description Get API information
	// @tags root
	// @produce json
	// @success 200 {object} string "Welcome to the URL Shortener API"
	// @router / [get]
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/urls", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/urls", func(c *gin.Context) {
		handler.GetShortUrls(c)
	})

	r.GET("/urls/:shortUrl", func(c *gin.Context) {
		handler.GetUrlObj(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":9808")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
