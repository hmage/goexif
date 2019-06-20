package goexif

import (
	"bytes"
	"testing"

	"github.com/hmage/goexif/exif"
	"github.com/hmage/goexif/mknote"
)

func TestHang1(t *testing.T) {
	exif.RegisterParsers(mknote.All...)

	input := []byte("II*\x00\t\x00\x00\x00\x00if\xcfe\t\x00\x02\x00\x00\xc0")
	result, err := exif.Decode(bytes.NewReader(input))
	if err.Error() != "exif: decode failed (EOF) " {
		t.Fatalf("expected resursive IFD error but got: %q", err.Error())
	}
	t.Log(result)
}

func TestRecursive1(t *testing.T) {
	exif.RegisterParsers(mknote.All...)

	input := []byte("MM\x00*\x00\x00\x00@\xbd\xbf\x01\x00\x00\x00\x00\xbd\xbf/\xbd\xbf" +
		"\xefY\xef\x02\x00\x00\x00\x01\x00\x00*\x00\x00\x00\x00\x00\x00\x00\x00 " +
		"\x00\x00\x00\x00\x00\x00\x00@\x00\x00\x00\x01*\x00\x00\x8d\b\x00\x00\x03" +
		"\x00\x00\x00*\x00\x00\x00\x00\x00*")
	result, err := exif.Decode(bytes.NewReader(input))
	if err.Error() != "exif: decode failed (tiff: recursive IFD) " {
		t.Fatalf("expected resursive IFD error but got: %q", err.Error())
	}
	t.Log(result)
}

func TestRecursive2(t *testing.T) {
	exif.RegisterParsers(mknote.All...)

	input := []byte("MM\x00*\x00\x00\x00\b\xd5\x02\x00\x00\x00\x19\xe3\xfd \x00\x00\x00" +
		"\x02\x00\x00\x00\x01՛\x00\x00\x00\b\xe3")
	result, err := exif.Decode(bytes.NewReader(input))
	if err.Error() != "exif: decode failed (tiff: recursive IFD) " {
		t.Fatalf("expected resursive IFD error but got: %q", err.Error())
	}
	t.Log(result)
}
