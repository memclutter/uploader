package uploader

const (
	CodeOk = iota
	CodeErrUnknownBucket
	CodeErrInvalidMimeType
	CodeErrMaxExceedFileSize
	CodeErrCreateDir
	CodeErrCreateFile
	CodeErrOpenFile
	CodeErrCopyFile
)
