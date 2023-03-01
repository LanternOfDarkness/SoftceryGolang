package storage

import (
	"os"
	"path/filepath"
	"strconv"
)

var Storage *ImageStorage
var LastImageID = 0

type ImageStorage struct {
	storagePath string
}

func NewStorage(storagePath string) *ImageStorage {
	return &ImageStorage{
		storagePath: storagePath,
	}
}


func GetNewImageID() int {
	return LastImageID + 1
}

func (fs *ImageStorage) LoadIdFromFiles() error {
	id := 0
	err := filepath.Walk(fs.storagePath + "100", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			id, err = strconv.Atoi(fileNameWithoutExt(info.Name()))
			if err != nil {
				return err
			}
		}
        return nil
    })
    if err != nil {
        return err
    }
	LastImageID = id
	return nil
}

func fileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}