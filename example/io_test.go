package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestCopyFunc(t *testing.T) {
	r := strings.NewReader("some io.reader stream to be read\n")

	io.Copy(os.Stdout, r)
}

func TestCopyBufferFunc(t *testing.T) {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 1)

	io.CopyBuffer(os.Stdout, r1, buf)
	io.CopyBuffer(os.Stdout, r2, buf)
}

func TestCopyNFunc(t *testing.T) {
	r := strings.NewReader("some io.reader stream to be read\n")

	io.CopyN(os.Stdout, r, 2)
}

func TestPipeFunc(t *testing.T) {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "hahaha")
		w.Close()
	}()

	io.Copy(os.Stdout, r)
}

func TestReadAtLeastFunc(t *testing.T) {
	r := strings.NewReader("some io.reader stream to be read\n")

	buf := make([]byte, 14)
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		fmt.Println("error is ", err)
	}
	fmt.Printf("%s\n", buf)

	shortbuf := make([]byte, 3)
	if _, err := io.ReadAtLeast(r, shortbuf, 4); err != nil {
		fmt.Println("error is: ", err)
	}
}

func TestWriteStringFunc(t *testing.T) {
	if _, err := io.WriteString(os.Stdout, "gogogo"); err != nil {
		fmt.Println("error is ", err)
	}
}

func TestReaderFunc(t *testing.T) {
	r := strings.NewReader("some io.reader stream to be read\n")
	lr := io.LimitReader(r, 3)
	io.Copy(os.Stdout, lr)
	fmt.Println()

	s := io.NewSectionReader(r, 5, 17)
	io.Copy(os.Stdout, s)
	fmt.Println()
	fmt.Printf("%d\n", s.Size())

	buf := make([]byte, 6)
	if _, err := s.ReadAt(buf, 10); err != nil {
		fmt.Println("error is", err)
	}
	fmt.Printf("%s\n", buf)

	if _, err := s.Seek(10, io.SeekStart); err != nil {
		fmt.Println("error is ", err)
	}
	if _, err := io.Copy(os.Stdout, s); err != nil {
		fmt.Println("error is", err)
	}
	fmt.Println()

	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	r3 := strings.NewReader("third reader\n")
	rr := io.MultiReader(r1, r2, r3)
	io.Copy(os.Stdout, rr)

	var buf1, buf2 strings.Builder
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		fmt.Println("error is ", err)
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}
