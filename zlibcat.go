package main

import (
	"compress/zlib"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	r, err := zlib.NewReader(file)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)
}

func usage() {
	println("zlibcat <filename>")
}
