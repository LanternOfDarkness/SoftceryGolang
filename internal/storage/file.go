package storage

import (
	"errors"
	"os"
	"path/filepath"
)



func (fs *ImageStorage) LoadImage(id string, quality string) ([]byte, error) {
  imgName := id + "_" + quality + ".jpg"
	imgPath := filepath.Join(fs.storagePath, imgName)
  if _, err := os.Stat(imgPath); os.IsNotExist(err) {
		return nil, errors.New("File not found")
	}

  image, err := os.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (fs *ImageStorage) SaveImage(quality string, image []byte) error {

  imgName := fs.getNewImageID() + "_" + quality + ".jpg"
  imgPath := filepath.Join(fs.storagePath, imgName)

  fo, err := os.Create(imgPath)
  if err != nil {
    return err
  }
  defer fo.Close()

  _, err = fo.Write(image)
  if err != nil {
    return err
  }
  
  return nil
}

func (fs *ImageStorage) getNewImageID() string {
  
  return "1"
}