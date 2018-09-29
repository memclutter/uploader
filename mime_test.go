package uploader

import (
	"strings"
	"testing"
)

var mimeTypesCheckTestTable = []struct {
	init MimeTypes
	in   string
	out  bool
}{
	{MimeTypes{"image/jpeg", "image/png"}, "image/jpeg", true},
	{MimeTypes{}, "image/png", true},
	{MimeTypes{"text/plain", "text/html"}, "image/jpeg", false},
}

func TestMimeTypes_Check(t *testing.T) {
	for _, tt := range mimeTypesCheckTestTable {
		t.Run(strings.Join(tt.init, ",")+tt.in, func(t *testing.T) {
			actual := tt.init.Check(tt.in)

			if actual != tt.out {
				t.Errorf("Incorrect result. Excepted %v, Actual %v", tt.out, actual)
			}
		})
	}
}
