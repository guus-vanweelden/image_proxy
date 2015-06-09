package cache

import (
	"errors"
	"image"
)

var imgCache map[string]image.Image

func init() {
	imgCache = make(map[string]image.Image)
}

func Get(key string) (image.Image, error) {
	if img, ok := imgCache[key]; ok {
		return img, nil
	}
	return nil, errors.New("")
}

func Set(key string, img image.Image) {
	imgCache[key] = img
}
