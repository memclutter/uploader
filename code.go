package uploader

const (
	CodeOk = iota
	CodeErrUnknownBucket
	CodeErrInvalidMimeType
	CodeErrDetectMimeType
	CodeErrMaxExceedFileSize
	CodeErrCreateDir
	CodeErrCreateFile
	CodeErrOpenFile
	CodeErrCopyFile
)
