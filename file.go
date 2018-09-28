package uploader

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"mime"
	"net/http"
)

func DetectContentType(file io.ReadSeeker) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	file.Seek(0, 0)

	return http.DetectContentType(buffer), nil
}

func GenerateRandomFileName(contentType string) (string, error) {
	name := GenerateRandomString(16)
	ext, err := DetectFileExtension(contentType)
	if err != nil {
		return "", err
	}

	return name + ext, nil
}

func GenerateRandomString(size int) string {
	buffer := make([]byte, size)
	rand.Read(buffer)
	return hex.EncodeToString(buffer)
}

func DetectFileExtension(contentType string) (string, error) {
	extensions, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", err
	}

	if len(extensions) == 0 {
		return "", errors.New("unknown file extension")
	}

	return extensions[0], nil
}
