package uploader

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// Represent uploader bucket.
type Bucket struct {
	Path        string
	BaseDir     string
	Dir         string
	MimeTypes   []string
	MaxFileSize int64
}

func (b *Bucket) Upload(multipartFile *multipart.FileHeader) Result {
	// TODO: Read real mime type
	mimeType := multipartFile.Header.Get("Content-Type")

	if !b.CheckMimeType(mimeType) {
		return Result{Code: CodeErrInvalidMimeType}
	}

	if multipartFile.Size > b.MaxFileSize {
		return Result{Code: CodeErrMaxExceedFileSize}
	}

	// TODO: generate filename (random or original)
	filename := multipartFile.Filename

	dirPath := path.Join(b.BaseDir, b.Dir)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			return Result{Code: CodeErrCreateDir}
		}
	}

	filePath := path.Join(dirPath, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return Result{Code: CodeErrCreateFile}
	}

	src, err := multipartFile.Open()
	if err != nil {
		return Result{Code: CodeErrOpenFile}
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		return Result{Code: CodeErrCopyFile}
	}

	return Result{
		Code:     CodeOk,
		Size:     multipartFile.Size,
		Name:     filename,
		MimeType: mimeType,
		Path:     path.Join(b.Path, filename),
	}
}

func (b *Bucket) CheckMimeType(mimeType string) bool {
	if len(b.MimeTypes) == 0 {
		return true
	}

	for _, mmt := range b.MimeTypes {
		if strings.Compare(mmt, mimeType) == 0 {
			return true
		}
	}

	return false
}
