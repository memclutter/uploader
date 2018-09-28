package uploader

const (
	CodeOk = iota
	CodeErrUnknownBucket
	CodeErrInvalidMimeType
	CodeErrDetectMimeType
	CodeErrExceedMaxFileSize
	CodeErrCreateDir
	CodeErrCreateFile
	CodeErrOpenFile
	CodeErrCopyFile
)
