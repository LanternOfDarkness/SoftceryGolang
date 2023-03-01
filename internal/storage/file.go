package storage

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)



func (fs *ImageStorage) LoadImage(id string, quality string) ([]byte, error) {
  imgName := id +  ".jpg"
	imgPath := filepath.Join(fs.storagePath + "/" + quality, imgName)
  if _, err := os.Stat(imgPath); os.IsNotExist(err) {
		return nil, errors.New("File not found")
	}

  image, err := os.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}

	return image, nil
}



func (fs *ImageStorage) SaveImage(quality string, img []byte) error {

  imgName := fmt.Sprint(GetNewImageID()) + ".jpg"
  imgPath := filepath.Join(fs.storagePath + "/" + quality, imgName)

  err := os.WriteFile(imgPath, img, 0644)
  if err != nil {
		log.Fatal(err)
    return err
  }

  return nil
}

