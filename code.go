package uploader

const (
	CodeOk = iota
	CodeErrUnknownBucket
	CodeErrInvalidMimeType
	CodeErrDetectMimeType
	CodeErrExceedMaxFileSize
	CodeErrGenerateFileName
	CodeErrCreateDir
	CodeErrCreateFile
	CodeErrOpenFile
	CodeErrCopyFile
)
