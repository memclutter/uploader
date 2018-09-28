package uploader

import (
	"io"
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
