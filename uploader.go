package uploader

import (
	"mime/multipart"
	"path"
)

type Uploader struct {
	baseDir string
	buckets map[string]Bucket
}

func NewUploader(baseDir string) *Uploader {
	return &Uploader{
		baseDir: baseDir,
		buckets: make(map[string]Bucket),
	}
}

func (u *Uploader) UploadFile(bucketName string, multipartFile *multipart.FileHeader) Result {
	if bucket, ok := u.buckets[bucketName]; !ok {
		return Result{Code: CodeErrUnknownBucket}
	} else {
		return bucket.Upload(multipartFile)
	}
}

func (u *Uploader) UploadForm(form *multipart.Form) (resultSet map[string][]Result) {
	resultSet = make(map[string][]Result)

	for bucketName, multipartFiles := range form.File {
		resultSet[bucketName] = make([]Result, len(multipartFiles))

		for i, multipartFile := range multipartFiles {
			resultSet[bucketName][i] = u.UploadFile(bucketName, multipartFile)
		}
	}

	return resultSet
}

// Register new uploader bucket
func (u *Uploader) RegisterBucket(name string, bucket Bucket) {
	if bucket.BaseDir == "" {
		bucket.BaseDir = u.baseDir
	}

	if bucket.Dir == "" {
		bucket.Dir = name
	}

	if bucket.Path == "" {
		bucket.Path = path.Join(bucket.Dir)
	}

	u.buckets[name] = bucket
}

// Unregister exists uploader bucket
func (u *Uploader) UnregisterBucket(name string) {
	delete(u.buckets, name)
}
