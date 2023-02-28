package storage

var Storage *ImageStorage

type ImageStorage struct {
	storagePath string
}

func NewStorage(storagePath string) *ImageStorage {
	return &ImageStorage{
		storagePath: storagePath,
	}
}
