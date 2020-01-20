package goexif

import (
	"bytes"
	"testing"

	"github.com/ZorinArsenij/goexif/exif"
	"github.com/ZorinArsenij/goexif/mknote"
)

func TestFuzzCrash1(t *testing.T) {
	exif.RegisterParsers(mknote.All...)

	input := []byte("II*\x00\t\x00\x00\x00\x00if\xcfe\t\x00\x02\x00\x00\xc0")
	result, err := exif.Decode(bytes.NewReader(input))
	if err.Error() != "exif: decode failed (invalid Count offset in tag) " {
		t.Fatalf("expected invalid count offset error but got: %q", err.Error())
	}
	t.Log(result)
}

func TestFuzzCrash2(t *testing.T) {
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

func TestFuzzCrash3(t *testing.T) {
	exif.RegisterParsers(mknote.All...)

	input := []byte("MM\x00*\x00\x00\x00\b\xd5\x02\x00\x00\x00\x19\xe3\xfd \x00\x00\x00" +
		"\x02\x00\x00\x00\x01՛\x00\x00\x00\b\xe3")
	result, err := exif.Decode(bytes.NewReader(input))
	if err.Error() != "exif: decode failed (tiff: recursive IFD) " {
		t.Fatalf("expected resursive IFD error but got: %q", err.Error())
	}
	t.Log(result)
}

func TestFuzzCrash4(t *testing.T) {
	exif.RegisterParsers(mknote.All...)

	input := []byte("II*\x00\b\x00\x00\x00\t\x000000000000" +
		"00000000000000000000" +
		"00000000000000000000" +
		"00000000000000000000" +
		"00000000000000000000" +
		"000000i\x87\x04\x00\x01\x00\x00\x00\xac\x00\x00\x0000" +
		"00000000000000000000" +
		"00000000000000000000" +
		"0000000000000000\x05\x00\x00\x00" +
		"\x00\xe00000")
	result, err := exif.Decode(bytes.NewReader(input))
	if err.Error() != "exif: decode failed (tiff: seek offset after EOF) " {
		t.Fatalf("expected seek offset after EOF error but got: %q", err.Error())
	}
	t.Log(result)
}
func TestFuzzCrash5(t *testing.T) {
	exif.RegisterParsers(mknote.All...)

	input := []byte("MM\x00*\x00\x00\x00\b\x00\a0000000000" +
		"00000000000000000000" +
		"000000000000000000\x87i" +
		"\x00\x04\x00\x00\x00\x0000000000000000" +
		"00000000000000000000" +
		"00000000000000")
	result, err := exif.Decode(bytes.NewReader(input))
	if err.Error() != "exif: decode failed (tiff: seek offset after EOF) " {
		t.Fatalf("expected seek offset after EOF error but got: %q", err.Error())
	}
	t.Log(result)
}
func TestFuzzCrash6(t *testing.T) {
	exif.RegisterParsers(mknote.All...)

	input := []byte("II*\x00\b\x00\x00\x000000\x05\x00\x00\x00\x00\xa000" +
		"00")
	result, err := exif.Decode(bytes.NewReader(input))
	if err.Error() != "exif: decode failed (invalid Count offset in tag) " {
		t.Fatalf("expected invalid Count offset error but got: %q", err.Error())
	}
	t.Log(result)
}
