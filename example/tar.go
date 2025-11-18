package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
)

func tarFunc() {
	fmt.Println("This is a placeholder for tar functionality.")

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	var files = []struct {
		Name, Body string
	}{
		{"file1.txt", "This is the content of file1."},
		{"file2.txt", "This is the content of file2."},
		{"readme.txt", "This is the readme file."},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			fmt.Println("Error writing header:", err)
			return
		}

		if _, err := tw.Write([]byte(file.Body)); err != nil {
			fmt.Println("Error writing body:", err)
			return
		}
	}

	if err := tw.Close(); err != nil {
		fmt.Println("Error closing tar writer:", err)
		return
	}

	fmt.Println("Tar archive created successfully")

	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error reading tar:", err)
			return
		}

		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			fmt.Println("Error copying file contents:", err)
			return
		}

		fmt.Println()
	}
}
