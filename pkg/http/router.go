package http

import (
	image_handler "github.com/LanternOfDarkness/SoftceryGolang/pkg/http/image"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/upload", uploadImageHandler)
	r.GET("/image/:id", downloadImageHandler)

	return r
}

func uploadImageHandler(c *gin.Context) {
  service := &image_handler.ImageUploadService{}
  service.UploadImage(c)
}

func downloadImageHandler(c *gin.Context) {
  service := &image_handler.ImageDownloadService{}
  service.DownloadImage(c)
}