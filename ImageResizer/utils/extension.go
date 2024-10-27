package utils

import (
	"errors"
	"mime/multipart"
	"net/http"
)

func ValidateImageExtension(src multipart.File) (string, error) {
	buffer := make([]byte, 512)
	if _, err := src.Read(buffer); err != nil {
		return "", errors.New("unable to read file for type detection")
	}

	contentType := http.DetectContentType(buffer)

	allowedTypes := map[string]string{
		"image/jpeg": ".jpeg",
		"image/png":  ".png",
		"image/jpg":  ".jpg",
	}

	if ext, exists := allowedTypes[contentType]; exists {
		return ext, nil
	}

	return "", errors.New("unsupported file type")
}
