package image_handler

import (
	"strconv"

	store "github.com/LanternOfDarkness/SoftceryGolang/internal/storage"
	"github.com/gin-gonic/gin"
)

type ImageDownloadService struct{}

func (s *ImageDownloadService) DownloadImage(c *gin.Context) {
	imageID := c.Param("id")
  quality := c.DefaultQuery("quality", "100")
  
  if validateQuality(quality) {

    image, err := store.Storage.LoadImage(imageID, quality)
    if err != nil {
      c.JSON(404, gin.H{ "message": "Image not found" })
      return
    }
    c.Data(200, "image/jpeg", image)
  } else {
    c.JSON(400, gin.H{ "message": "Quality is invalid. Valid values are 100, 75, 50, 25." })
  }

}

func validateQuality(quality string) bool {
  var qualityInt int
  var err error
  if qualityInt, err = strconv.Atoi(quality); err != nil {
    return false
  }
  var qualityList = []int{100, 75, 50, 25}
  for _, q := range qualityList {
    if q == qualityInt {
      return true
    }
  }
  return false
}