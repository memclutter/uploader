package uploader

const (
	CodeOk = iota
	CodeErrDetectMimeType
	CodeErrInvalidMimeType
	CodeErrExceedMaxFileSize
	CodeErrGenerateFileName
	CodeErrCreateDir
	CodeErrCreateFile
	CodeErrOpenFile
	CodeErrCopyFile
)
