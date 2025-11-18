package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func byteFunc() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Fprintf(&b, "world!\n")
	b.WriteTo(os.Stdout)

	b.Write([]byte("abcde"))
	fmt.Printf("%s\n", b.Next(2))
	fmt.Printf("%s\n", b.Next(2))
	fmt.Printf("%s\n", b.Next(2))

	b.Write([]byte("qwert"))
	a := make([]byte, 1)
	b.Read(a)
	fmt.Printf("%s\n", a)
	fmt.Println(b.String())

	buf := bytes.NewBufferString("aaaaaaaaaa\n")
	io.Copy(os.Stdout, buf)

	before, after, found := bytes.Cut([]byte("gopher"), []byte("ph"))
	fmt.Println(string(before), string(after), found)
}
