package main

import (
	"log"
	"os"
	"testing"
)

func TestReadFileFunc(t *testing.T) {
	data, err := os.ReadFile("../file/testdata.txt")
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(data)
}

func TestWriteFileFunc(t *testing.T) {
	err := os.WriteFile("../file/testwrite.txt", []byte("gjjjjjjjjjjjjjjj"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
