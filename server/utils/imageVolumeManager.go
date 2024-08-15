package utils

import (
	"os"
)

func SaveImageToImagesVolume(image []byte, imageName string) error {
	// Get volume path
	volumePath := os.Getenv("DATA_VOLUME_PATH")
	if volumePath == "" {
		panic("Failed to get volume path")
	}
	// Save image to volume
	return os.WriteFile(volumePath+"/"+imageName+".png", image, 0644)
}

func GetImageFromImagesVolume(imageName string) ([]byte, error) {
	// Get volume path
	volumePath := os.Getenv("DATA_VOLUME_PATH")
	if volumePath == "" {
		panic("Failed to get volume path")
	}
	// Read image from volume
	return os.ReadFile(volumePath + "/" + imageName + ".png")
}
