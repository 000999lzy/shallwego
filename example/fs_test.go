package main

import (
	"fmt"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestReadFile(t *testing.T) {
	fsys := fstest.MapFS{
		"hello.txt": {
			Data: []byte("hello gogogo!\n"),
		},
	}

	data, err := fs.ReadFile(fsys, "hello.txt")
	if err != nil {
		fmt.Println("errors")
	}

	fmt.Print(string(data))
}
