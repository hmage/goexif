#!/usr/bin/env bash

go-fuzz-build github.com/hmage/goexif
ulimit -m 1048576
ulimit -v 1048576
go-fuzz -bin=goexif-fuzz.zip
