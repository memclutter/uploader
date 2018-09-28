package uploader

type Result struct {
	Code     int
	Path     string
	Name     string
	MimeType string
	Size     int64
}
