package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

func zipFunc() {
	readZipFile("../file/testzip.zip")
	writeZipFile()
}

func writeZipFile() {
	buf := new(bytes.Buffer)

	zw := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := zw.Create(file.Name)
		if err != nil {
			fmt.Println("Error creating file in zip:", err)
			return
		}

		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println("Error writing to file in zip:", err)
			return
		}
	}

	err := zw.Close()
	if err != nil {
		fmt.Println("Error closing zip writer:", err)
		return
	}
}

func readZipFile(filePath string) {
	r, err := zip.OpenReader(filePath)
	if err != nil {
		fmt.Println("Error opening zip:", err)
		return
	}

	defer r.Close()

	for _, f := range r.File {
		fmt.Println("File in zip:", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}

		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
		}

		rc.Close()
		fmt.Println()
	}
}
