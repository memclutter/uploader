package uploader

const (
	CodeOk = iota
	CodeErrInvalidMimeType
	CodeErrDetectMimeType
	CodeErrExceedMaxFileSize
	CodeErrGenerateFileName
	CodeErrCreateDir
	CodeErrCreateFile
	CodeErrOpenFile
	CodeErrCopyFile
)
