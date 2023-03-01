package image_handler

import (
	"net/http"

	queue "github.com/LanternOfDarkness/SoftceryGolang/internal/queue"
	store "github.com/LanternOfDarkness/SoftceryGolang/internal/storage"
	"github.com/gin-gonic/gin"
)

type ImageUploadService struct{}

func (s *ImageUploadService) UploadImage(c *gin.Context) {
	
  formFile, err := c.FormFile("image")
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  file, err := formFile.Open()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  defer file.Close()

  imgBytes := make([]byte, formFile.Size)
  _, err = file.Read(imgBytes)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  if err := queue.SendImageToQueue(imgBytes); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Image uploaded successfully.",
    "imageId": store.GetNewImageID(),
	})
}


