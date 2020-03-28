package main

import (
	"compress/zlib"
	"flag"
	"io"
	"os"
)

func main() {
	var filename = flag.String("filepath", "", "file to unzip")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	r, err := zlib.NewReader(file)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)
}
