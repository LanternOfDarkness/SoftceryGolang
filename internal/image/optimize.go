package image

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"

	store "github.com/LanternOfDarkness/SoftceryGolang/internal/storage"
	"github.com/nfnt/resize"
)

type ImageSize struct {
	Width  int
	Height int
}

var qualityList = []int{75, 50, 25}

func ProcessImage(img []byte) error {
	err := store.Storage.SaveImage("100", img)
	if err != nil {
		return err
	}
	for _, quality := range qualityList {
		err := ResizeImage(img, quality); if err != nil {
			return err
		}
	}
	store.LastImageID = store.GetNewImageID()
	return nil
}

func ResizeImage(file []byte, quality int) error {
	var newWidth, newHeight uint
	
	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return err
	}
	imgSize := GetImageSize(img)
	newWidth = uint(imgSize.Width * quality / 100)
	newHeight = uint(imgSize.Height * quality / 100)

	newImage := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	bytesBuff := new(bytes.Buffer)

	if err := jpeg.Encode(bytesBuff, newImage, nil); err != nil {
		return err
	}
	imgBytes := bytesBuff.Bytes()

	if err := store.Storage.SaveImage(fmt.Sprint(quality), imgBytes); err != nil {
		return err
	}

	return nil
}

func GetImageSize(img image.Image) (ImageSize) {
	var imgSize ImageSize
	imgSize.Width = img.Bounds().Max.X
	imgSize.Height = img.Bounds().Max.Y

	return imgSize
}