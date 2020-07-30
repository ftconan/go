// Author: magician
// File:   main.go
// Date:   2020/7/30
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" // register PNG decoder
	"io"
	"os"
)

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
