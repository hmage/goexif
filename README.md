EXIF container parser in pure Go
======

[![GoDoc](https://godoc.org/github.com/ZorinArsenij/goexif?status.svg)](https://godoc.org/github.com/ZorinArsenij/goexif) ![CI](https://action-badges.now.sh/ZorinArsenij/goexif)

A fork of [rwcarlsen/goexif](https://github.com/rwcarlsen/goexif) that fixes bugs found via fuzzing to prevent infinite loops or OOMs when invalid or badly formed exif input is fed into it.

Suggestions and pull requests are welcome.

Example usage:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ZorinArsenij/goexif/exif"
	"github.com/ZorinArsenij/goexif/mknote"
)

func ExampleDecode() {
	// Optionally register camera makenote data parsing - currently Nikon and
	// Canon are supported.
	exif.RegisterParsers(mknote.All...)

	fname := "sample1.jpg"

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
	fmt.Println(camModel.StringVal())

	focal, _ := x.Get(exif.FocalLength)
	numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
	fmt.Printf("%v/%v", numer, denom)

	// Two convenience functions exist for date/time taken and GPS coords:
	tm, _ := x.DateTime()
	fmt.Println("Taken: ", tm)

	lat, long, _ := x.LatLong()
	fmt.Println("lat, long: ", lat, ", ", long)
}
```

<!--golang-->
