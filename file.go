package uploader

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"mime"
	"net/http"
)

func detectContentType(file io.ReadSeeker) (string, error) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	file.Seek(0, 0)

	return http.DetectContentType(buffer), nil
}

func generateRandomFileName(contentType string) (string, error) {
	name := generateRandomString(16)
	ext, err := detectFileExtension(contentType)
	if err != nil {
		return "", err
	}

	return name + ext, nil
}

func generateRandomString(size int) string {
	buffer := make([]byte, size)
	rand.Read(buffer)
	return hex.EncodeToString(buffer)
}

func detectFileExtension(contentType string) (string, error) {
	extensions, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", err
	}

	if len(extensions) == 0 {
		return "", errors.New("unknown file extension")
	}

	return extensions[0], nil
}
