package goexif

import (
	"bytes"

	"github.com/hmage/goexif/exif"
	"github.com/hmage/goexif/mknote"
)

func Fuzz(data []byte) int {
	exif.RegisterParsers(mknote.All...)

	if _, err := exif.Decode(bytes.NewReader(data)); err != nil {
		return 0
	}
	return 1
}
