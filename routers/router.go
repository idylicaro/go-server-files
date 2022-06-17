package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idylicaro/go-server-files/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/static", "./public")
	apiv1 := router.Group("/api/v1")
	{

		apiv1.GET("ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		apiv1.POST("/upload", controllers.UploadFile)

	}
	return router
}
