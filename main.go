package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idylicaro/go-server-files/controllers"
)

func main() {
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/static", "./public")
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("welcome", func(c *gin.Context) {
				c.String(http.StatusOK, "Welcome")
			})

			v1.POST("/upload", controllers.UploadFile)
		}
	}
	router.Run(":8080")
}
