package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"time"
)

func compressFunc() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	zw.Name = "example.txt"
	zw.Comment = "This is an example of gzip compression in Go."
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)
	data := []byte("Hello, World! Hello, World! Hello, World! Hello, World!")
	_, err := zw.Write(data)
	if err != nil {
		fmt.Println("Error writing data to gzip writer:", err)
		return
	}

	if err := zw.Close(); err != nil {
		fmt.Println("Error closing gzip writer:", err)
		return
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}

	fmt.Println("Gzip File Name:", zr.Name)
	fmt.Println("Gzip Comment:", zr.Comment)
	fmt.Println("Gzip ModTime:", zr.ModTime)

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		fmt.Println("Error reading from gzip reader:", err)
		return
	}

	if err := zr.Close(); err != nil {
		fmt.Println("Error closing gzip reader:", err)
		return
	}

}
