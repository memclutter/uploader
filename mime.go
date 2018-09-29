package uploader

import "strings"

type MimeTypes []string

func (mt MimeTypes) Check(mimeType string) bool {

	// allow all mime types
	if len(mt) == 0 {
		return true
	}

	for _, t := range mt {
		if strings.Compare(t, mimeType) == 0 {
			return true
		}
	}

	return false
}
