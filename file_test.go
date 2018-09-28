package uploader

import (
	"strings"
	"testing"
)

var detectFileExtensionTestTable = []struct {
	in     string
	out    string
	hasErr bool
}{
	{"image/jpeg", ".jpg", false},
	{"image/png", ".png", false},
}

func TestDetectFileExtension(t *testing.T) {
	for _, tt := range detectFileExtensionTestTable {
		t.Run(tt.in, func(t *testing.T) {
			actual, err := detectFileExtension(tt.in)

			if tt.hasErr {
				if err == nil {
					t.Errorf("No error however it is expected")
				}
			} else {
				if strings.Compare(actual, tt.out) != 0 {
					t.Errorf("Incorrect result. Excepted %s, Actual %s", tt.out, actual)
				}
			}
		})
	}
}
