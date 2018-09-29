package uploader

import (
	"io"
	"mime/multipart"
	"os"
	"path"
)

// Represent uploader bucket.
type Bucket struct {
	Path        string
	BaseDir     string
	Dir         string
	MimeTypes   MimeTypes
	MaxFileSize int64
}

func (b *Bucket) upload(multipartFile *multipart.FileHeader) Result {
	if multipartFile.Size > b.MaxFileSize {
		return Result{Code: CodeErrExceedMaxFileSize}
	}

	src, err := multipartFile.Open()
	if err != nil {
		return Result{Code: CodeErrOpenFile}
	}

	mimeType, err := detectContentType(src)
	if err != nil {
		return Result{Code: CodeErrDetectMimeType}
	}

	if !b.MimeTypes.Check(mimeType) {
		return Result{Code: CodeErrInvalidMimeType}
	}

	filename, err := generateRandomFileName(mimeType)
	if err != nil {
		return Result{Code: CodeErrGenerateFileName}
	}

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
