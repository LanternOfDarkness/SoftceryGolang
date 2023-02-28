package image_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageUploadService struct{}

func (s *ImageUploadService) UploadImage(c *gin.Context) {
	
  image, err := c.FormFile("image")
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := c.SaveUploadedFile(image, "./images/1_100.jpg"); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }



	c.JSON(http.StatusAccepted, gin.H{
		"message": "Image uploaded successfully",
	})
}
